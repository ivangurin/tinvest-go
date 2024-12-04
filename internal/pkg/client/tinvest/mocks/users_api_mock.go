// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"
	contractv1 "tinvest-go/internal/pb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// UsersAPIMock is an autogenerated mock type for the UsersAPIMock type
type UsersAPIMock struct {
	mock.Mock
}

type UsersAPIMock_Expecter struct {
	mock *mock.Mock
}

func (_m *UsersAPIMock) EXPECT() *UsersAPIMock_Expecter {
	return &UsersAPIMock_Expecter{mock: &_m.Mock}
}

// GetAccounts provides a mock function with given fields: ctx, in, opts
func (_m *UsersAPIMock) GetAccounts(ctx context.Context, in *contractv1.GetAccountsRequest, opts ...grpc.CallOption) (*contractv1.GetAccountsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAccounts")
	}

	var r0 *contractv1.GetAccountsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetAccountsRequest, ...grpc.CallOption) (*contractv1.GetAccountsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetAccountsRequest, ...grpc.CallOption) *contractv1.GetAccountsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractv1.GetAccountsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contractv1.GetAccountsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersAPIMock_GetAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccounts'
type UsersAPIMock_GetAccounts_Call struct {
	*mock.Call
}

// GetAccounts is a helper method to define mock.On call
//   - ctx context.Context
//   - in *contractv1.GetAccountsRequest
//   - opts ...grpc.CallOption
func (_e *UsersAPIMock_Expecter) GetAccounts(ctx interface{}, in interface{}, opts ...interface{}) *UsersAPIMock_GetAccounts_Call {
	return &UsersAPIMock_GetAccounts_Call{Call: _e.mock.On("GetAccounts",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *UsersAPIMock_GetAccounts_Call) Run(run func(ctx context.Context, in *contractv1.GetAccountsRequest, opts ...grpc.CallOption)) *UsersAPIMock_GetAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*contractv1.GetAccountsRequest), variadicArgs...)
	})
	return _c
}

func (_c *UsersAPIMock_GetAccounts_Call) Return(_a0 *contractv1.GetAccountsResponse, _a1 error) *UsersAPIMock_GetAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UsersAPIMock_GetAccounts_Call) RunAndReturn(run func(context.Context, *contractv1.GetAccountsRequest, ...grpc.CallOption) (*contractv1.GetAccountsResponse, error)) *UsersAPIMock_GetAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// GetInfo provides a mock function with given fields: ctx, in, opts
func (_m *UsersAPIMock) GetInfo(ctx context.Context, in *contractv1.GetInfoRequest, opts ...grpc.CallOption) (*contractv1.GetInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInfo")
	}

	var r0 *contractv1.GetInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetInfoRequest, ...grpc.CallOption) (*contractv1.GetInfoResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetInfoRequest, ...grpc.CallOption) *contractv1.GetInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractv1.GetInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contractv1.GetInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersAPIMock_GetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInfo'
type UsersAPIMock_GetInfo_Call struct {
	*mock.Call
}

// GetInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - in *contractv1.GetInfoRequest
//   - opts ...grpc.CallOption
func (_e *UsersAPIMock_Expecter) GetInfo(ctx interface{}, in interface{}, opts ...interface{}) *UsersAPIMock_GetInfo_Call {
	return &UsersAPIMock_GetInfo_Call{Call: _e.mock.On("GetInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *UsersAPIMock_GetInfo_Call) Run(run func(ctx context.Context, in *contractv1.GetInfoRequest, opts ...grpc.CallOption)) *UsersAPIMock_GetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*contractv1.GetInfoRequest), variadicArgs...)
	})
	return _c
}

func (_c *UsersAPIMock_GetInfo_Call) Return(_a0 *contractv1.GetInfoResponse, _a1 error) *UsersAPIMock_GetInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UsersAPIMock_GetInfo_Call) RunAndReturn(run func(context.Context, *contractv1.GetInfoRequest, ...grpc.CallOption) (*contractv1.GetInfoResponse, error)) *UsersAPIMock_GetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetMarginAttributes provides a mock function with given fields: ctx, in, opts
func (_m *UsersAPIMock) GetMarginAttributes(ctx context.Context, in *contractv1.GetMarginAttributesRequest, opts ...grpc.CallOption) (*contractv1.GetMarginAttributesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMarginAttributes")
	}

	var r0 *contractv1.GetMarginAttributesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetMarginAttributesRequest, ...grpc.CallOption) (*contractv1.GetMarginAttributesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetMarginAttributesRequest, ...grpc.CallOption) *contractv1.GetMarginAttributesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractv1.GetMarginAttributesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contractv1.GetMarginAttributesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersAPIMock_GetMarginAttributes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMarginAttributes'
type UsersAPIMock_GetMarginAttributes_Call struct {
	*mock.Call
}

// GetMarginAttributes is a helper method to define mock.On call
//   - ctx context.Context
//   - in *contractv1.GetMarginAttributesRequest
//   - opts ...grpc.CallOption
func (_e *UsersAPIMock_Expecter) GetMarginAttributes(ctx interface{}, in interface{}, opts ...interface{}) *UsersAPIMock_GetMarginAttributes_Call {
	return &UsersAPIMock_GetMarginAttributes_Call{Call: _e.mock.On("GetMarginAttributes",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *UsersAPIMock_GetMarginAttributes_Call) Run(run func(ctx context.Context, in *contractv1.GetMarginAttributesRequest, opts ...grpc.CallOption)) *UsersAPIMock_GetMarginAttributes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*contractv1.GetMarginAttributesRequest), variadicArgs...)
	})
	return _c
}

func (_c *UsersAPIMock_GetMarginAttributes_Call) Return(_a0 *contractv1.GetMarginAttributesResponse, _a1 error) *UsersAPIMock_GetMarginAttributes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UsersAPIMock_GetMarginAttributes_Call) RunAndReturn(run func(context.Context, *contractv1.GetMarginAttributesRequest, ...grpc.CallOption) (*contractv1.GetMarginAttributesResponse, error)) *UsersAPIMock_GetMarginAttributes_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserTariff provides a mock function with given fields: ctx, in, opts
func (_m *UsersAPIMock) GetUserTariff(ctx context.Context, in *contractv1.GetUserTariffRequest, opts ...grpc.CallOption) (*contractv1.GetUserTariffResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetUserTariff")
	}

	var r0 *contractv1.GetUserTariffResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetUserTariffRequest, ...grpc.CallOption) (*contractv1.GetUserTariffResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contractv1.GetUserTariffRequest, ...grpc.CallOption) *contractv1.GetUserTariffResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractv1.GetUserTariffResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contractv1.GetUserTariffRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersAPIMock_GetUserTariff_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserTariff'
type UsersAPIMock_GetUserTariff_Call struct {
	*mock.Call
}

// GetUserTariff is a helper method to define mock.On call
//   - ctx context.Context
//   - in *contractv1.GetUserTariffRequest
//   - opts ...grpc.CallOption
func (_e *UsersAPIMock_Expecter) GetUserTariff(ctx interface{}, in interface{}, opts ...interface{}) *UsersAPIMock_GetUserTariff_Call {
	return &UsersAPIMock_GetUserTariff_Call{Call: _e.mock.On("GetUserTariff",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *UsersAPIMock_GetUserTariff_Call) Run(run func(ctx context.Context, in *contractv1.GetUserTariffRequest, opts ...grpc.CallOption)) *UsersAPIMock_GetUserTariff_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*contractv1.GetUserTariffRequest), variadicArgs...)
	})
	return _c
}

func (_c *UsersAPIMock_GetUserTariff_Call) Return(_a0 *contractv1.GetUserTariffResponse, _a1 error) *UsersAPIMock_GetUserTariff_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UsersAPIMock_GetUserTariff_Call) RunAndReturn(run func(context.Context, *contractv1.GetUserTariffRequest, ...grpc.CallOption) (*contractv1.GetUserTariffResponse, error)) *UsersAPIMock_GetUserTariff_Call {
	_c.Call.Return(run)
	return _c
}

// NewUsersAPIMock creates a new instance of UsersAPIMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersAPIMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersAPIMock {
	mock := &UsersAPIMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
