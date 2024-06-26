// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "applicationDesignTest/internal/model"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// availabilityRepository is an autogenerated mock type for the availabilityRepository type
type availabilityRepository struct {
	mock.Mock
}

// DecrementQuotaForRange provides a mock function with given fields: from, to
func (_m *availabilityRepository) DecrementQuotaForRange(from time.Time, to time.Time) {
	_m.Called(from, to)
}

// GetAll provides a mock function with given fields:
func (_m *availabilityRepository) GetAll() []model.RoomAvailability {
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
func (_m *availabilityRepository) Replace(index int, record model.RoomAvailability) error {
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

// newAvailabilityRepository creates a new instance of availabilityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newAvailabilityRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *availabilityRepository {
	mock := &availabilityRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
