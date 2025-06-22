// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package events

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"

	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"
)

// EventClient wraps Kubernetes event recording functionality for ACK controllers
type EventClient struct {
	// k8sClient is the Kubernetes client for creating events
	k8sClient kubernetes.Interface
	// recorder is the event recorder for legacy v1 events
	recorder record.EventRecorder
	// controllerName is the name of the ACK controller emitting events
	controllerName string
	// instanceID is a unique identifier for this controller instance
	instanceID string
}

// NewEventClient creates a new EventClient for ACK controllers
func NewEventClient(
	k8sClient kubernetes.Interface,
	recorder record.EventRecorder,
	controllerName string,
	instanceID string,
) *EventClient {
	return &EventClient{
		k8sClient:      k8sClient,
		recorder:       recorder,
		controllerName: controllerName,
		instanceID:     instanceID,
	}
}

// EmitEvent creates a Kubernetes event for the given resource
func (c *EventClient) EmitEvent(
	ctx context.Context,
	resource acktypes.AWSResource,
	eventType EventType,
	reason EventReason,
	message string,
) error {
	metaObj := resource.MetaObject()
	runtimeObj := resource.RuntimeObject()
	
	// Use the new events/v1 API for better structured events
	event := &eventsv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.generateEventName(resource, reason),
			Namespace: metaObj.GetNamespace(),
		},
		Regarding: corev1.ObjectReference{
			APIVersion: runtimeObj.GetObjectKind().GroupVersionKind().GroupVersion().String(),
			Kind:       runtimeObj.GetObjectKind().GroupVersionKind().Kind,
			Name:       metaObj.GetName(),
			Namespace:  metaObj.GetNamespace(),
			UID:        metaObj.GetUID(),
		},
		Type:                string(eventType),
		Reason:              string(reason),
		Note:                message,
		Action:              c.getActionForReason(reason),
		EventTime:           metav1.MicroTime{Time: time.Now()},
		ReportingController: c.controllerName,
		ReportingInstance:   c.instanceID,
	}

	_, err := c.k8sClient.EventsV1().Events(metaObj.GetNamespace()).Create(ctx, event, metav1.CreateOptions{})
	return err
}

// EmitNormalEvent emits a Normal type event
func (c *EventClient) EmitNormalEvent(
	ctx context.Context,
	resource acktypes.AWSResource,
	reason EventReason,
	message string,
) error {
	return c.EmitEvent(ctx, resource, EventTypeNormal, reason, message)
}

// EmitWarningEvent emits a Warning type event
func (c *EventClient) EmitWarningEvent(
	ctx context.Context,
	resource acktypes.AWSResource,
	reason EventReason,
	message string,
) error {
	return c.EmitEvent(ctx, resource, EventTypeWarning, reason, message)
}

// EmitEventForObject creates a Kubernetes event for any metav1.Object
func (c *EventClient) EmitEventForObject(
	ctx context.Context,
	obj metav1.Object,
	eventType EventType,
	reason EventReason,
	message string,
) error {
	// Get the runtime object to access ObjectKind
	runtimeObj, ok := obj.(runtime.Object)
	if !ok {
		// Fall back to basic implementation without GVK information
		return c.emitBasicEventForObject(ctx, obj, eventType, reason, message)
	}
	
	// Use the new events/v1 API for better structured events
	event := &eventsv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.generateEventNameForObject(obj, reason),
			Namespace: obj.GetNamespace(),
		},
		Regarding: corev1.ObjectReference{
			APIVersion: runtimeObj.GetObjectKind().GroupVersionKind().GroupVersion().String(),
			Kind:       runtimeObj.GetObjectKind().GroupVersionKind().Kind,
			Name:       obj.GetName(),
			Namespace:  obj.GetNamespace(),
			UID:        obj.GetUID(),
		},
		Type:                string(eventType),
		Reason:              string(reason),
		Note:                message,
		Action:              c.getActionForReason(reason),
		EventTime:           metav1.MicroTime{Time: time.Now()},
		ReportingController: c.controllerName,
		ReportingInstance:   c.instanceID,
	}

	_, err := c.k8sClient.EventsV1().Events(obj.GetNamespace()).Create(ctx, event, metav1.CreateOptions{})
	return err
}

// emitBasicEventForObject creates a basic event without GVK information
func (c *EventClient) emitBasicEventForObject(
	ctx context.Context,
	obj metav1.Object,
	eventType EventType,
	reason EventReason,
	message string,
) error {
	// Try to cast to runtime.Object for the event recorder
	if runtimeObj, ok := obj.(runtime.Object); ok {
		c.recorder.Event(runtimeObj, string(eventType), string(reason), message)
	}
	// If it's not a runtime.Object, we can't emit the event via the recorder
	// but we don't return an error as this is a fallback scenario
	return nil
}

// EmitNormalEventForObject emits a Normal type event for metav1.Object
func (c *EventClient) EmitNormalEventForObject(
	ctx context.Context,
	obj metav1.Object,
	reason EventReason,
	message string,
) error {
	return c.EmitEventForObject(ctx, obj, EventTypeNormal, reason, message)
}

// EmitWarningEventForObject emits a Warning type event for metav1.Object
func (c *EventClient) EmitWarningEventForObject(
	ctx context.Context,
	obj metav1.Object,
	reason EventReason,
	message string,
) error {
	return c.EmitEventForObject(ctx, obj, EventTypeWarning, reason, message)
}

// generateEventName generates a unique name for the event
func (c *EventClient) generateEventName(resource acktypes.AWSResource, reason EventReason) string {
	// Use a combination of resource name, reason, and timestamp to ensure uniqueness
	timestamp := time.Now().UnixNano()
	metaObj := resource.MetaObject()
	return fmt.Sprintf("%s-%s-%s-%d", metaObj.GetName(), string(reason), c.instanceID, timestamp)
}

// generateEventNameForObject generates a unique name for the event for metav1.Object
func (c *EventClient) generateEventNameForObject(obj metav1.Object, reason EventReason) string {
	// Use a combination of resource name, reason, and timestamp to ensure uniqueness
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s-%s-%s-%d", obj.GetName(), string(reason), c.instanceID, timestamp)
}

// getActionForReason maps event reasons to actions
func (c *EventClient) getActionForReason(reason EventReason) string {
	switch reason {
	case EventReasonResourceCreated, EventReasonResourceCreating:
		return "Creating"
	case EventReasonResourceUpdated, EventReasonResourceUpdating:
		return "Updating"
	case EventReasonResourceDeleted, EventReasonResourceDeleting:
		return "Deleting"
	case EventReasonResourceAdopted, EventReasonResourceAdopting:
		return "Adopting"
	case EventReasonResourceSynced:
		return "Syncing"
	case EventReasonResourceAvailable:
		return "Available"
	case EventReasonResourceTerminal:
		return "Terminal"
	case EventReasonAPIError, EventReasonPermissionError, EventReasonValidationError:
		return "Error"
	default:
		return "Unknown"
	}
}