// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	crypto "github.com/fat/common/crypto"
	mock "github.com/stretchr/testify/mock"

	models "github.com/fat/models"
)

// IndodaxWebhook is an autogenerated mock type for the IndodaxWebhook type
type IndodaxWebhook struct {
	mock.Mock
}

// GetInfo provides a mock function with given fields: payload
func (_m *IndodaxWebhook) GetInfo(payload *crypto.PayloadSHA512) (*models.ResponseInfoIndodax, error) {
	ret := _m.Called(payload)

	var r0 *models.ResponseInfoIndodax
	if rf, ok := ret.Get(0).(func(*crypto.PayloadSHA512) *models.ResponseInfoIndodax); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseInfoIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*crypto.PayloadSHA512) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: payload
func (_m *IndodaxWebhook) GetOrder(payload *crypto.PayloadSHA512) (*models.ResponseGetOrderIndodax, error) {
	ret := _m.Called(payload)

	var r0 *models.ResponseGetOrderIndodax
	if rf, ok := ret.Get(0).(func(*crypto.PayloadSHA512) *models.ResponseGetOrderIndodax); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseGetOrderIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*crypto.PayloadSHA512) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderHistory provides a mock function with given fields: payload
func (_m *IndodaxWebhook) GetOrderHistory(payload *crypto.PayloadSHA512) (*models.ResponseOrderHistoryIndodax, error) {
	ret := _m.Called(payload)

	var r0 *models.ResponseOrderHistoryIndodax
	if rf, ok := ret.Get(0).(func(*crypto.PayloadSHA512) *models.ResponseOrderHistoryIndodax); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseOrderHistoryIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*crypto.PayloadSHA512) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkePendingtHistory provides a mock function with given fields: symbol
func (_m *IndodaxWebhook) MarkePendingtHistory(symbol string) (*models.ResponseMarketPendingIndodax, error) {
	ret := _m.Called(symbol)

	var r0 *models.ResponseMarketPendingIndodax
	if rf, ok := ret.Get(0).(func(string) *models.ResponseMarketPendingIndodax); ok {
		r0 = rf(symbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseMarketPendingIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketHistory provides a mock function with given fields: payload
func (_m *IndodaxWebhook) MarketHistory(payload *models.MarketHistoryPayload) ([]models.MarketHistoryIndodax, error) {
	ret := _m.Called(payload)

	var r0 []models.MarketHistoryIndodax
	if rf, ok := ret.Get(0).(func(*models.MarketHistoryPayload) []models.MarketHistoryIndodax); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.MarketHistoryIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.MarketHistoryPayload) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Order provides a mock function with given fields: payload
func (_m *IndodaxWebhook) Order(payload *crypto.PayloadSHA512) (*models.ResponseOrderIndodax, error) {
	ret := _m.Called(payload)

	var r0 *models.ResponseOrderIndodax
	if rf, ok := ret.Get(0).(func(*crypto.PayloadSHA512) *models.ResponseOrderIndodax); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseOrderIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*crypto.PayloadSHA512) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
