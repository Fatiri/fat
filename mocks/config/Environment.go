// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	models "github.com/fat/models"
	mock "github.com/stretchr/testify/mock"
)

// Environment is an autogenerated mock type for the Environment type
type Environment struct {
	mock.Mock
}

// InitEnvironment provides a mock function with given fields:
func (_m *Environment) InitEnvironment() (*models.Environment, error) {
	ret := _m.Called()

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func() *models.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
