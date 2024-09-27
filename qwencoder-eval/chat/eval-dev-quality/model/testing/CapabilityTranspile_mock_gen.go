// Code generated by mockery v2.40.3. DO NOT EDIT.

package modeltesting

import (
	mock "github.com/stretchr/testify/mock"
	metrics "github.com/symflower/eval-dev-quality/evaluate/metrics"

	model "github.com/symflower/eval-dev-quality/model"
)

// MockCapabilityTranspile is an autogenerated mock type for the CapabilityTranspile type
type MockCapabilityTranspile struct {
	mock.Mock
}

// Transpile provides a mock function with given fields: ctx
func (_m *MockCapabilityTranspile) Transpile(ctx model.Context) (metrics.Assessments, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Transpile")
	}

	var r0 metrics.Assessments
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Context) (metrics.Assessments, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(model.Context) metrics.Assessments); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Assessments)
		}
	}

	if rf, ok := ret.Get(1).(func(model.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockCapabilityTranspile creates a new instance of MockCapabilityTranspile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCapabilityTranspile(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCapabilityTranspile {
	mock := &MockCapabilityTranspile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
