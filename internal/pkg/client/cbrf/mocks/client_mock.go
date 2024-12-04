// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// ClientMock is an autogenerated mock type for the ClientMock type
type ClientMock struct {
	mock.Mock
}

type ClientMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientMock) EXPECT() *ClientMock_Expecter {
	return &ClientMock_Expecter{mock: &_m.Mock}
}

// GetExchangeRate provides a mock function with given fields: ctx, currencyID, date
func (_m *ClientMock) GetExchangeRate(ctx context.Context, currencyID string, date time.Time) (float64, error) {
	ret := _m.Called(ctx, currencyID, date)

	if len(ret) == 0 {
		panic("no return value specified for GetExchangeRate")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) (float64, error)); ok {
		return rf(ctx, currencyID, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) float64); ok {
		r0 = rf(ctx, currencyID, date)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, currencyID, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientMock_GetExchangeRate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExchangeRate'
type ClientMock_GetExchangeRate_Call struct {
	*mock.Call
}

// GetExchangeRate is a helper method to define mock.On call
//   - ctx context.Context
//   - currencyID string
//   - date time.Time
func (_e *ClientMock_Expecter) GetExchangeRate(ctx interface{}, currencyID interface{}, date interface{}) *ClientMock_GetExchangeRate_Call {
	return &ClientMock_GetExchangeRate_Call{Call: _e.mock.On("GetExchangeRate", ctx, currencyID, date)}
}

func (_c *ClientMock_GetExchangeRate_Call) Run(run func(ctx context.Context, currencyID string, date time.Time)) *ClientMock_GetExchangeRate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Time))
	})
	return _c
}

func (_c *ClientMock_GetExchangeRate_Call) Return(_a0 float64, _a1 error) *ClientMock_GetExchangeRate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientMock_GetExchangeRate_Call) RunAndReturn(run func(context.Context, string, time.Time) (float64, error)) *ClientMock_GetExchangeRate_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientMock creates a new instance of ClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientMock {
	mock := &ClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}