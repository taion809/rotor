// Code generated by MockGen. DO NOT EDIT.
// Source: fromflags.go

package http

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFromFlags is a mock of FromFlags interface
type MockFromFlags struct {
	ctrl     *gomock.Controller
	recorder *MockFromFlagsMockRecorder
}

// MockFromFlagsMockRecorder is the mock recorder for MockFromFlags
type MockFromFlagsMockRecorder struct {
	mock *MockFromFlags
}

// NewMockFromFlags creates a new mock instance
func NewMockFromFlags(ctrl *gomock.Controller) *MockFromFlags {
	mock := &MockFromFlags{ctrl: ctrl}
	mock.recorder = &MockFromFlagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFromFlags) EXPECT() *MockFromFlagsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockFromFlags) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockFromFlagsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockFromFlags)(nil).Validate))
}

// MakeEndpoint mocks base method
func (m *MockFromFlags) MakeEndpoint() (Endpoint, error) {
	ret := m.ctrl.Call(m, "MakeEndpoint")
	ret0, _ := ret[0].(Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeEndpoint indicates an expected call of MakeEndpoint
func (mr *MockFromFlagsMockRecorder) MakeEndpoint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeEndpoint", reflect.TypeOf((*MockFromFlags)(nil).MakeEndpoint))
}
