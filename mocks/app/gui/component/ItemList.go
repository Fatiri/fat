// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fyne "fyne.io/fyne/v2"
	component "github.com/fat/app/gui/component"

	mock "github.com/stretchr/testify/mock"

	widget "fyne.io/fyne/v2/widget"
)

// ItemList is an autogenerated mock type for the ItemList type
type ItemList struct {
	mock.Mock
}

// V1 provides a mock function with given fields: apps, content
func (_m *ItemList) V1(apps []component.AppInfo, content *fyne.Container) *widget.List {
	ret := _m.Called(apps, content)

	var r0 *widget.List
	if rf, ok := ret.Get(0).(func([]component.AppInfo, *fyne.Container) *widget.List); ok {
		r0 = rf(apps, content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.List)
		}
	}

	return r0
}

// V2 provides a mock function with given fields: items
func (_m *ItemList) V2(items *[]string) *widget.List {
	ret := _m.Called(items)

	var r0 *widget.List
	if rf, ok := ret.Get(0).(func(*[]string) *widget.List); ok {
		r0 = rf(items)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widget.List)
		}
	}

	return r0
}
