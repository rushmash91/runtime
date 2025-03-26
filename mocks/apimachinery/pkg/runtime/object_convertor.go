// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ObjectConvertor is an autogenerated mock type for the ObjectConvertor type
type ObjectConvertor struct {
	mock.Mock
}

// Convert provides a mock function with given fields: in, out, context
func (_m *ObjectConvertor) Convert(in interface{}, out interface{}, context interface{}) error {
	ret := _m.Called(in, out, context)

	if len(ret) == 0 {
		panic("no return value specified for Convert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, interface{}) error); ok {
		r0 = rf(in, out, context)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConvertFieldLabel provides a mock function with given fields: gvk, label, value
func (_m *ObjectConvertor) ConvertFieldLabel(gvk schema.GroupVersionKind, label string, value string) (string, string, error) {
	ret := _m.Called(gvk, label, value)

	if len(ret) == 0 {
		panic("no return value specified for ConvertFieldLabel")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(schema.GroupVersionKind, string, string) (string, string, error)); ok {
		return rf(gvk, label, value)
	}
	if rf, ok := ret.Get(0).(func(schema.GroupVersionKind, string, string) string); ok {
		r0 = rf(gvk, label, value)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(schema.GroupVersionKind, string, string) string); ok {
		r1 = rf(gvk, label, value)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(schema.GroupVersionKind, string, string) error); ok {
		r2 = rf(gvk, label, value)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ConvertToVersion provides a mock function with given fields: in, gv
func (_m *ObjectConvertor) ConvertToVersion(in runtime.Object, gv runtime.GroupVersioner) (runtime.Object, error) {
	ret := _m.Called(in, gv)

	if len(ret) == 0 {
		panic("no return value specified for ConvertToVersion")
	}

	var r0 runtime.Object
	var r1 error
	if rf, ok := ret.Get(0).(func(runtime.Object, runtime.GroupVersioner) (runtime.Object, error)); ok {
		return rf(in, gv)
	}
	if rf, ok := ret.Get(0).(func(runtime.Object, runtime.GroupVersioner) runtime.Object); ok {
		r0 = rf(in, gv)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Object)
		}
	}

	if rf, ok := ret.Get(1).(func(runtime.Object, runtime.GroupVersioner) error); ok {
		r1 = rf(in, gv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewObjectConvertor creates a new instance of ObjectConvertor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObjectConvertor(t interface {
	mock.TestingT
	Cleanup(func())
}) *ObjectConvertor {
	mock := &ObjectConvertor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
