// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/gateway/throttler (interfaces: Throttler)

// Package mocks_gateway is a generated GoMock package.
package mocks_gateway

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockThrottler is a mock of Throttler interface.
type MockThrottler struct {
	ctrl     *gomock.Controller
	recorder *MockThrottlerMockRecorder
}

// MockThrottlerMockRecorder is the mock recorder for MockThrottler.
type MockThrottlerMockRecorder struct {
	mock *MockThrottler
}

// NewMockThrottler creates a new mock instance.
func NewMockThrottler(ctrl *gomock.Controller) *MockThrottler {
	mock := &MockThrottler{ctrl: ctrl}
	mock.recorder = &MockThrottlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThrottler) EXPECT() *MockThrottlerMockRecorder {
	return m.recorder
}

// CheckLimitReached mocks base method.
func (m *MockThrottler) CheckLimitReached(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckLimitReached", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckLimitReached indicates an expected call of CheckLimitReached.
func (mr *MockThrottlerMockRecorder) CheckLimitReached(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLimitReached", reflect.TypeOf((*MockThrottler)(nil).CheckLimitReached), arg0, arg1)
}
