// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// LogWriterStrategy is an autogenerated mock type for the LogWriterStrategy type
type LogWriterStrategy struct {
	mock.Mock
}

// Write provides a mock function with given fields: ctx, reader
func (_m *LogWriterStrategy) Write(ctx context.Context, reader io.Reader) <-chan error {
	ret := _m.Called(ctx, reader)

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader) <-chan error); ok {
		r0 = rf(ctx, reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

type mockConstructorTestingTNewLogWriterStrategy interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogWriterStrategy creates a new instance of LogWriterStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogWriterStrategy(t mockConstructorTestingTNewLogWriterStrategy) *LogWriterStrategy {
	mock := &LogWriterStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
