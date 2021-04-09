// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/masterraf21/reksti-ordering-backend/models"
	mock "github.com/stretchr/testify/mock"
)

// OrderDetailsRepository is an autogenerated mock type for the OrderDetailsRepository type
type OrderDetailsRepository struct {
	mock.Mock
}

// BulkInsert provides a mock function with given fields: ctx, orderdetails
func (_m *OrderDetailsRepository) BulkInsert(ctx context.Context, orderdetails []models.OrderDetails) error {
	ret := _m.Called(ctx, orderdetails)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []models.OrderDetails) error); ok {
		r0 = rf(ctx, orderdetails)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, orderDetailsID
func (_m *OrderDetailsRepository) DeleteByID(ctx context.Context, orderDetailsID uint32) error {
	ret := _m.Called(ctx, orderDetailsID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, orderDetailsID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *OrderDetailsRepository) GetAll() ([]models.OrderDetails, error) {
	ret := _m.Called()

	var r0 []models.OrderDetails
	if rf, ok := ret.Get(0).(func() []models.OrderDetails); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.OrderDetails)
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

// GetByID provides a mock function with given fields: orderDetailsID
func (_m *OrderDetailsRepository) GetByID(orderDetailsID uint32) (*models.OrderDetails, error) {
	ret := _m.Called(orderDetailsID)

	var r0 *models.OrderDetails
	if rf, ok := ret.Get(0).(func(uint32) *models.OrderDetails); ok {
		r0 = rf(orderDetailsID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrderDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(orderDetailsID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderDetailsByOrderID provides a mock function with given fields: orderID
func (_m *OrderDetailsRepository) GetOrderDetailsByOrderID(orderID uint32) ([]models.OrderDetails, error) {
	ret := _m.Called(orderID)

	var r0 []models.OrderDetails
	if rf, ok := ret.Get(0).(func(uint32) []models.OrderDetails); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.OrderDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, ord
func (_m *OrderDetailsRepository) Store(ctx context.Context, ord *models.OrderDetails) (uint32, error) {
	ret := _m.Called(ctx, ord)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, *models.OrderDetails) uint32); ok {
		r0 = rf(ctx, ord)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.OrderDetails) error); ok {
		r1 = rf(ctx, ord)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateArbitrary provides a mock function with given fields: ctx, orderDetailsID, columnName, value
func (_m *OrderDetailsRepository) UpdateArbitrary(ctx context.Context, orderDetailsID uint32, columnName string, value interface{}) error {
	ret := _m.Called(ctx, orderDetailsID, columnName, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, string, interface{}) error); ok {
		r0 = rf(ctx, orderDetailsID, columnName, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateByID provides a mock function with given fields: ctx, orderDetailsID, order
func (_m *OrderDetailsRepository) UpdateByID(ctx context.Context, orderDetailsID uint32, order *models.OrderDetails) error {
	ret := _m.Called(ctx, orderDetailsID, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, *models.OrderDetails) error); ok {
		r0 = rf(ctx, orderDetailsID, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTotalPrice provides a mock function with given fields: ctx, orderDetailsID
func (_m *OrderDetailsRepository) UpdateTotalPrice(ctx context.Context, orderDetailsID uint32) error {
	ret := _m.Called(ctx, orderDetailsID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, orderDetailsID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
