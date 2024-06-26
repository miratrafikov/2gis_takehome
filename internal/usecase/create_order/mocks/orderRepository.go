// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "applicationDesignTest/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// orderRepository is an autogenerated mock type for the orderRepository type
type orderRepository struct {
	mock.Mock
}

// Push provides a mock function with given fields: order
func (_m *orderRepository) Push(order model.Order) []model.Order {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for Push")
	}

	var r0 []model.Order
	if rf, ok := ret.Get(0).(func(model.Order) []model.Order); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Order)
		}
	}

	return r0
}

// newOrderRepository creates a new instance of orderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *orderRepository {
	mock := &orderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
