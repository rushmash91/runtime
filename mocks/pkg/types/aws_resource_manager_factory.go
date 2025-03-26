// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	config "github.com/aws-controllers-k8s/runtime/pkg/config"
	aws "github.com/aws/aws-sdk-go-v2/aws"
	logr "github.com/go-logr/logr"

	metrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"

	mock "github.com/stretchr/testify/mock"

	types "github.com/aws-controllers-k8s/runtime/pkg/types"

	v1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// AWSResourceManagerFactory is an autogenerated mock type for the AWSResourceManagerFactory type
type AWSResourceManagerFactory struct {
	mock.Mock
}

// IsAdoptable provides a mock function with no fields
func (_m *AWSResourceManagerFactory) IsAdoptable() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAdoptable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ManagerFor provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7
func (_m *AWSResourceManagerFactory) ManagerFor(_a0 config.Config, _a1 aws.Config, _a2 logr.Logger, _a3 *metrics.Metrics, _a4 types.Reconciler, _a5 v1alpha1.AWSAccountID, _a6 v1alpha1.AWSRegion, _a7 v1alpha1.AWSResourceName) (types.AWSResourceManager, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)

	if len(ret) == 0 {
		panic("no return value specified for ManagerFor")
	}

	var r0 types.AWSResourceManager
	var r1 error
	if rf, ok := ret.Get(0).(func(config.Config, aws.Config, logr.Logger, *metrics.Metrics, types.Reconciler, v1alpha1.AWSAccountID, v1alpha1.AWSRegion, v1alpha1.AWSResourceName) (types.AWSResourceManager, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	}
	if rf, ok := ret.Get(0).(func(config.Config, aws.Config, logr.Logger, *metrics.Metrics, types.Reconciler, v1alpha1.AWSAccountID, v1alpha1.AWSRegion, v1alpha1.AWSResourceName) types.AWSResourceManager); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResourceManager)
		}
	}

	if rf, ok := ret.Get(1).(func(config.Config, aws.Config, logr.Logger, *metrics.Metrics, types.Reconciler, v1alpha1.AWSAccountID, v1alpha1.AWSRegion, v1alpha1.AWSResourceName) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequeueOnSuccessSeconds provides a mock function with no fields
func (_m *AWSResourceManagerFactory) RequeueOnSuccessSeconds() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequeueOnSuccessSeconds")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ResourceDescriptor provides a mock function with no fields
func (_m *AWSResourceManagerFactory) ResourceDescriptor() types.AWSResourceDescriptor {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ResourceDescriptor")
	}

	var r0 types.AWSResourceDescriptor
	if rf, ok := ret.Get(0).(func() types.AWSResourceDescriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResourceDescriptor)
		}
	}

	return r0
}

// NewAWSResourceManagerFactory creates a new instance of AWSResourceManagerFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAWSResourceManagerFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *AWSResourceManagerFactory {
	mock := &AWSResourceManagerFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
