// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	product "lendra/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductServiceInterface is an autogenerated mock type for the ProductServiceInterface type
type ProductServiceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: userIdLogin, input
func (_m *ProductServiceInterface) Create(userIdLogin int, input product.Product) error {
	ret := _m.Called(userIdLogin, input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, product.Product) error); ok {
		r0 = rf(userIdLogin, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: userIdLogin, idProduct
func (_m *ProductServiceInterface) Delete(userIdLogin int, idProduct int) error {
	ret := _m.Called(userIdLogin, idProduct)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(userIdLogin, idProduct)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProduct provides a mock function with given fields:
func (_m *ProductServiceInterface) GetAllProduct() ([]product.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllProduct")
	}

	var r0 []product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]product.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []product.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductById provides a mock function with given fields: idProduct
func (_m *ProductServiceInterface) GetProductById(idProduct int) (*product.Product, error) {
	ret := _m.Called(idProduct)

	if len(ret) == 0 {
		panic("no return value specified for GetProductById")
	}

	var r0 *product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*product.Product, error)); ok {
		return rf(idProduct)
	}
	if rf, ok := ret.Get(0).(func(int) *product.Product); ok {
		r0 = rf(idProduct)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userIdLogin, idProduct, input
func (_m *ProductServiceInterface) Update(userIdLogin int, idProduct int, input product.Product) error {
	ret := _m.Called(userIdLogin, idProduct, input)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, product.Product) error); ok {
		r0 = rf(userIdLogin, idProduct, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductServiceInterface creates a new instance of ProductServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductServiceInterface {
	mock := &ProductServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
