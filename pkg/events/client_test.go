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
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/record"
)

func TestNewEventClient(t *testing.T) {
	// Create a fake Kubernetes client
	fakeClient := fake.NewSimpleClientset()
	
	// Create a fake event recorder
	fakeRecorder := record.NewFakeRecorder(10)
	
	// Create event client
	eventClient := NewEventClient(fakeClient, fakeRecorder, "test-controller", "test-instance")
	
	assert.NotNil(t, eventClient)
	assert.Equal(t, "test-controller", eventClient.controllerName)
	assert.Equal(t, "test-instance", eventClient.instanceID)
}

func TestFormatMessage(t *testing.T) {
	tests := []struct {
		name         string
		baseMessage  string
		details      []string
		expected     string
	}{
		{
			name:        "No details",
			baseMessage: "Resource created",
			details:     []string{},
			expected:    "Resource created",
		},
		{
			name:        "Single detail",
			baseMessage: "Resource created",
			details:     []string{"successfully"},
			expected:    "Resource created: successfully",
		},
		{
			name:        "Multiple details",
			baseMessage: "Resource created",
			details:     []string{"successfully", "with ID xyz123"},
			expected:    "Resource created: successfully: with ID xyz123",
		},
		{
			name:        "Empty detail",
			baseMessage: "Resource created",
			details:     []string{""},
			expected:    "Resource created",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatMessage(tt.baseMessage, tt.details...)
			assert.Equal(t, tt.expected, result)
		})
	}
}