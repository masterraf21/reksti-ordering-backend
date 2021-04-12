// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/masterraf21/reksti-ordering-backend/models"
	mock "github.com/stretchr/testify/mock"
)

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: ctx, CustomerID
func (_m *CustomerRepository) DeleteByID(ctx context.Context, CustomerID uint32) error {
	ret := _m.Called(ctx, CustomerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, CustomerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *CustomerRepository) GetAll() ([]models.Customer, error) {
	ret := _m.Called()

	var r0 []models.Customer
	if rf, ok := ret.Get(0).(func() []models.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Customer)
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

// GetByID provides a mock function with given fields: CustomerID
func (_m *CustomerRepository) GetByID(CustomerID uint32) (*models.Customer, error) {
	ret := _m.Called(CustomerID)

	var r0 *models.Customer
	if rf, ok := ret.Get(0).(func(uint32) *models.Customer); ok {
		r0 = rf(CustomerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(CustomerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, cust
func (_m *CustomerRepository) Store(ctx context.Context, cust *models.Customer) (uint32, error) {
	ret := _m.Called(ctx, cust)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, *models.Customer) uint32); ok {
		r0 = rf(ctx, cust)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Customer) error); ok {
		r1 = rf(ctx, cust)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
