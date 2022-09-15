// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	product "project/e-commerce/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductData is an autogenerated mock type for the DataInterface type
type ProductData struct {
	mock.Mock
}

// DeleteByToken provides a mock function with given fields: param, token
func (_m *ProductData) DeleteByToken(param int, token int) (int, error) {
	ret := _m.Called(param, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(param, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(param, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: data
func (_m *ProductData) InsertData(data product.Core) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(product.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(product.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllProduct provides a mock function with given fields: page, category
func (_m *ProductData) SelectAllProduct(page int, category string) ([]product.Core, error) {
	ret := _m.Called(page, category)

	var r0 []product.Core
	if rf, ok := ret.Get(0).(func(int, string) []product.Core); ok {
		r0 = rf(page, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(page, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: param
func (_m *ProductData) SelectById(param int) (product.Core, error) {
	ret := _m.Called(param)

	var r0 product.Core
	if rf, ok := ret.Get(0).(func(int) product.Core); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(product.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectMyProduct provides a mock function with given fields: token
func (_m *ProductData) SelectMyProduct(token int) ([]product.Core, error) {
	ret := _m.Called(token)

	var r0 []product.Core
	if rf, ok := ret.Get(0).(func(int) []product.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateData provides a mock function with given fields: token, newData
func (_m *ProductData) UpdateData(token int, newData product.Core) (int, error) {
	ret := _m.Called(token, newData)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, product.Core) int); ok {
		r0 = rf(token, newData)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, product.Core) error); ok {
		r1 = rf(token, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductData interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductData creates a new instance of ProductData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductData(t mockConstructorTestingTNewProductData) *ProductData {
	mock := &ProductData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
