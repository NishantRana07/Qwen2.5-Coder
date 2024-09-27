// Code generated by mockery v2.40.3. DO NOT EDIT.

package providertesting

import (
	mock "github.com/stretchr/testify/mock"
	log "github.com/symflower/eval-dev-quality/log"

	model "github.com/symflower/eval-dev-quality/model"
)

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

// Available provides a mock function with given fields: logger
func (_m *MockProvider) Available(logger *log.Logger) error {
	ret := _m.Called(logger)

	if len(ret) == 0 {
		panic("no return value specified for Available")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*log.Logger) error); ok {
		r0 = rf(logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *MockProvider) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Models provides a mock function with given fields:
func (_m *MockProvider) Models() ([]model.Model, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Models")
	}

	var r0 []model.Model
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Model, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Model); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Model)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockProvider creates a new instance of MockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProvider {
	mock := &MockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
