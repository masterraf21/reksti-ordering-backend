// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/masterraf21/reksti-ordering-backend/models"
	mock "github.com/stretchr/testify/mock"
)

// MenuRepository is an autogenerated mock type for the MenuRepository type
type MenuRepository struct {
	mock.Mock
}

// BulkInsert provides a mock function with given fields: ctx, Menu
func (_m *MenuRepository) BulkInsert(ctx context.Context, Menu []models.Menu) error {
	ret := _m.Called(ctx, Menu)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []models.Menu) error); ok {
		r0 = rf(ctx, Menu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, menuID
func (_m *MenuRepository) DeleteByID(ctx context.Context, menuID uint32) error {
	ret := _m.Called(ctx, menuID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, menuID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *MenuRepository) GetAll() ([]models.Menu, error) {
	ret := _m.Called()

	var r0 []models.Menu
	if rf, ok := ret.Get(0).(func() []models.Menu); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Menu)
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

// GetByID provides a mock function with given fields: menuID
func (_m *MenuRepository) GetByID(menuID uint32) (*models.Menu, error) {
	ret := _m.Called(menuID)

	var r0 *models.Menu
	if rf, ok := ret.Get(0).(func(uint32) *models.Menu); ok {
		r0 = rf(menuID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(menuID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, ord
func (_m *MenuRepository) Store(ctx context.Context, ord *models.Menu) (uint32, error) {
	ret := _m.Called(ctx, ord)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, *models.Menu) uint32); ok {
		r0 = rf(ctx, ord)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Menu) error); ok {
		r1 = rf(ctx, ord)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, menuID, order
func (_m *MenuRepository) UpdateByID(ctx context.Context, menuID uint32, order *models.Menu) error {
	ret := _m.Called(ctx, menuID, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, *models.Menu) error); ok {
		r0 = rf(ctx, menuID, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
