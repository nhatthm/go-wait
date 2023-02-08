// Code generated by mockery v2.18.0. DO NOT EDIT.

package wait

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Waiter is an autogenerated mock type for the Waiter type
type Waiter struct {
	mock.Mock
}

// Wait provides a mock function with given fields: ctx
func (_m *Waiter) Wait(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewWaiter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWaiter creates a new instance of Waiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWaiter(t mockConstructorTestingTNewWaiter) *Waiter {
	mock := &Waiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
