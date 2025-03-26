// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	meta "k8s.io/apimachinery/pkg/api/meta"

	mock "github.com/stretchr/testify/mock"

	runtime "k8s.io/apimachinery/pkg/runtime"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	types "k8s.io/apimachinery/pkg/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, obj, opts
func (_m *Client) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, ...client.CreateOption) error); ok {
		r0 = rf(ctx, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, obj, opts
func (_m *Client) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, ...client.DeleteOption) error); ok {
		r0 = rf(ctx, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllOf provides a mock function with given fields: ctx, obj, opts
func (_m *Client) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllOf")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, ...client.DeleteAllOfOption) error); ok {
		r0 = rf(ctx, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key, obj, opts
func (_m *Client) Get(ctx context.Context, key types.NamespacedName, obj client.Object, opts ...client.GetOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.NamespacedName, client.Object, ...client.GetOption) error); ok {
		r0 = rf(ctx, key, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupVersionKindFor provides a mock function with given fields: obj
func (_m *Client) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for GroupVersionKindFor")
	}

	var r0 schema.GroupVersionKind
	var r1 error
	if rf, ok := ret.Get(0).(func(runtime.Object) (schema.GroupVersionKind, error)); ok {
		return rf(obj)
	}
	if rf, ok := ret.Get(0).(func(runtime.Object) schema.GroupVersionKind); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Get(0).(schema.GroupVersionKind)
	}

	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsObjectNamespaced provides a mock function with given fields: obj
func (_m *Client) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for IsObjectNamespaced")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(runtime.Object) (bool, error)); ok {
		return rf(obj)
	}
	if rf, ok := ret.Get(0).(func(runtime.Object) bool); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, list, opts
func (_m *Client) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, list)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ObjectList, ...client.ListOption) error); ok {
		r0 = rf(ctx, list, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Patch provides a mock function with given fields: ctx, obj, patch, opts
func (_m *Client) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj, patch)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Patch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, client.Patch, ...client.PatchOption) error); ok {
		r0 = rf(ctx, obj, patch, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RESTMapper provides a mock function with no fields
func (_m *Client) RESTMapper() meta.RESTMapper {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RESTMapper")
	}

	var r0 meta.RESTMapper
	if rf, ok := ret.Get(0).(func() meta.RESTMapper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(meta.RESTMapper)
		}
	}

	return r0
}

// Scheme provides a mock function with no fields
func (_m *Client) Scheme() *runtime.Scheme {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Scheme")
	}

	var r0 *runtime.Scheme
	if rf, ok := ret.Get(0).(func() *runtime.Scheme); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.Scheme)
		}
	}

	return r0
}

// Status provides a mock function with no fields
func (_m *Client) Status() client.SubResourceWriter {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 client.SubResourceWriter
	if rf, ok := ret.Get(0).(func() client.SubResourceWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SubResourceWriter)
		}
	}

	return r0
}

// SubResource provides a mock function with given fields: subResource
func (_m *Client) SubResource(subResource string) client.SubResourceClient {
	ret := _m.Called(subResource)

	if len(ret) == 0 {
		panic("no return value specified for SubResource")
	}

	var r0 client.SubResourceClient
	if rf, ok := ret.Get(0).(func(string) client.SubResourceClient); ok {
		r0 = rf(subResource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SubResourceClient)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx, obj, opts
func (_m *Client) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, ...client.UpdateOption) error); ok {
		r0 = rf(ctx, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
