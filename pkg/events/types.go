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

// EventType represents the type of Kubernetes event
type EventType string

const (
	// EventTypeNormal represents normal events
	EventTypeNormal EventType = "Normal"
	// EventTypeWarning represents warning events
	EventTypeWarning EventType = "Warning"
)

// EventReason represents standardized reasons for ACK events
type EventReason string

const (
	// Resource lifecycle events
	EventReasonResourceCreating  EventReason = "ResourceCreating"
	EventReasonResourceCreated   EventReason = "ResourceCreated"
	EventReasonResourceUpdating  EventReason = "ResourceUpdating"
	EventReasonResourceUpdated   EventReason = "ResourceUpdated"
	EventReasonResourceDeleting  EventReason = "ResourceDeleting"
	EventReasonResourceDeleted   EventReason = "ResourceDeleted"
	EventReasonResourceSynced    EventReason = "ResourceSynced"
	EventReasonResourceAvailable EventReason = "ResourceAvailable"
	EventReasonResourceTerminal  EventReason = "ResourceTerminal"

	// Adoption events
	EventReasonResourceAdopting EventReason = "ResourceAdopting"
	EventReasonResourceAdopted  EventReason = "ResourceAdopted"

	// Field export events
	EventReasonFieldExportStarted   EventReason = "FieldExportStarted"
	EventReasonFieldExportCompleted EventReason = "FieldExportCompleted"
	EventReasonFieldExportFailed    EventReason = "FieldExportFailed"

	// Error events
	EventReasonAPIError        EventReason = "APIError"
	EventReasonPermissionError EventReason = "PermissionError"
	EventReasonValidationError EventReason = "ValidationError"
	EventReasonThrottlingError EventReason = "ThrottlingError"
	EventReasonReconcileError  EventReason = "ReconcileError"

	// State transition events
	EventReasonStateTransition EventReason = "StateTransition"
)

// EventMessage provides standardized event messages for common ACK operations
type EventMessage struct {
	// Resource lifecycle messages
	ResourceCreating  string
	ResourceCreated   string
	ResourceUpdating  string
	ResourceUpdated   string
	ResourceDeleting  string
	ResourceDeleted   string
	ResourceSynced    string
	ResourceAvailable string
	ResourceTerminal  string

	// Adoption messages
	ResourceAdopting string
	ResourceAdopted  string

	// Field export messages
	FieldExportStarted   string  
	FieldExportCompleted string
	FieldExportFailed    string

	// Error messages
	APIError        string
	PermissionError string
	ValidationError string
	ThrottlingError string
	ReconcileError  string

	// State transition messages
	StateTransition string
}

// StandardMessages provides default event messages for ACK operations
var StandardMessages = EventMessage{
	// Resource lifecycle messages
	ResourceCreating:  "Resource creation initiated",
	ResourceCreated:   "Resource created successfully",
	ResourceUpdating:  "Resource update initiated", 
	ResourceUpdated:   "Resource updated successfully",
	ResourceDeleting:  "Resource deletion initiated",
	ResourceDeleted:   "Resource deleted successfully",
	ResourceSynced:    "Resource synchronized with AWS",
	ResourceAvailable: "Resource is available and ready",
	ResourceTerminal:  "Resource is in terminal state",

	// Adoption messages
	ResourceAdopting: "Resource adoption initiated",
	ResourceAdopted:  "Resource adopted successfully",

	// Field export messages  
	FieldExportStarted:   "Field export operation started",
	FieldExportCompleted: "Field export operation completed",
	FieldExportFailed:    "Field export operation failed",

	// Error messages
	APIError:        "AWS API error occurred",
	PermissionError: "Permission denied for AWS operation",
	ValidationError: "Resource validation failed",
	ThrottlingError: "AWS API request was throttled",
	ReconcileError:  "Resource reconciliation failed",

	// State transition messages
	StateTransition: "Resource state changed",
}

// FormatMessage formats an event message with additional context
func FormatMessage(baseMessage string, details ...string) string {
	message := baseMessage
	for _, detail := range details {
		if detail != "" {
			message += ": " + detail
		}
	}
	return message
}