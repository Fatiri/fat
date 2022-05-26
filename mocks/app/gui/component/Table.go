// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	component "github.com/fat/app/gui/component"
	mock "github.com/stretchr/testify/mock"

	widget "fyne.io/fyne/v2/widget"
)

// Table is an autogenerated mock type for the Table type
type Table struct {
	mock.Mock
}

// V1 provides a mock function with given fields: data
func (_m *Table) V1(data *component.List) *widget.Table {
	ret := _m.Called(data)

	var r0 *widget.Table
	if rf, ok := ret.Get(0).(func(*component.List) *widget.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.Table)
		}
	}

	return r0
}

// V2 provides a mock function with given fields: data
func (_m *Table) V2(data *component.List) *widget.Table {
	ret := _m.Called(data)

	var r0 *widget.Table
	if rf, ok := ret.Get(0).(func(*component.List) *widget.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.Table)
		}
	}

	return r0
}

// V3 provides a mock function with given fields: data
func (_m *Table) V3(data *component.List) *widget.Table {
	ret := _m.Called(data)

	var r0 *widget.Table
	if rf, ok := ret.Get(0).(func(*component.List) *widget.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.Table)
		}
	}

	return r0
}

// V4 provides a mock function with given fields: data
func (_m *Table) V4(data *component.List) *widget.Table {
	ret := _m.Called(data)

	var r0 *widget.Table
	if rf, ok := ret.Get(0).(func(*component.List) *widget.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.Table)
		}
	}

	return r0
}

// V5 provides a mock function with given fields: data
func (_m *Table) V5(data *component.List) *widget.Table {
	ret := _m.Called(data)

	var r0 *widget.Table
	if rf, ok := ret.Get(0).(func(*component.List) *widget.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.Table)
		}
	}

	return r0
}
