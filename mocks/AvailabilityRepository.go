// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "applicationDesignTest/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// AvailabilityRepository is an autogenerated mock type for the AvailabilityRepository type
type AvailabilityRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *AvailabilityRepository) GetAll() []model.RoomAvailability {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []model.RoomAvailability
	if rf, ok := ret.Get(0).(func() []model.RoomAvailability); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.RoomAvailability)
		}
	}

	return r0
}

// Replace provides a mock function with given fields: index, record
func (_m *AvailabilityRepository) Replace(index int, record model.RoomAvailability) error {
	ret := _m.Called(index, record)

	if len(ret) == 0 {
		panic("no return value specified for Replace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, model.RoomAvailability) error); ok {
		r0 = rf(index, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAvailabilityRepository creates a new instance of AvailabilityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAvailabilityRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AvailabilityRepository {
	mock := &AvailabilityRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
