// Code generated by MockGen. DO NOT EDIT.
// Source: api_config.go

package flags

import (
	gomock "github.com/golang/mock/gomock"
	http "github.com/turbinelabs/api/http"
	reflect "reflect"
)

// MockAPIConfigFromFlags is a mock of APIConfigFromFlags interface
type MockAPIConfigFromFlags struct {
	ctrl     *gomock.Controller
	recorder *MockAPIConfigFromFlagsMockRecorder
}

// MockAPIConfigFromFlagsMockRecorder is the mock recorder for MockAPIConfigFromFlags
type MockAPIConfigFromFlagsMockRecorder struct {
	mock *MockAPIConfigFromFlags
}

// NewMockAPIConfigFromFlags creates a new mock instance
func NewMockAPIConfigFromFlags(ctrl *gomock.Controller) *MockAPIConfigFromFlags {
	mock := &MockAPIConfigFromFlags{ctrl: ctrl}
	mock.recorder = &MockAPIConfigFromFlagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIConfigFromFlags) EXPECT() *MockAPIConfigFromFlagsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockAPIConfigFromFlags) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockAPIConfigFromFlagsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAPIConfigFromFlags)(nil).Validate))
}

// MakeEndpoint mocks base method
func (m *MockAPIConfigFromFlags) MakeEndpoint() (http.Endpoint, error) {
	ret := m.ctrl.Call(m, "MakeEndpoint")
	ret0, _ := ret[0].(http.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeEndpoint indicates an expected call of MakeEndpoint
func (mr *MockAPIConfigFromFlagsMockRecorder) MakeEndpoint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeEndpoint", reflect.TypeOf((*MockAPIConfigFromFlags)(nil).MakeEndpoint))
}

// APIKey mocks base method
func (m *MockAPIConfigFromFlags) APIKey() string {
	ret := m.ctrl.Call(m, "APIKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// APIKey indicates an expected call of APIKey
func (mr *MockAPIConfigFromFlagsMockRecorder) APIKey() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIKey", reflect.TypeOf((*MockAPIConfigFromFlags)(nil).APIKey))
}

// APIAuthKeyFromFlags mocks base method
func (m *MockAPIConfigFromFlags) APIAuthKeyFromFlags() APIAuthKeyFromFlags {
	ret := m.ctrl.Call(m, "APIAuthKeyFromFlags")
	ret0, _ := ret[0].(APIAuthKeyFromFlags)
	return ret0
}

// APIAuthKeyFromFlags indicates an expected call of APIAuthKeyFromFlags
func (mr *MockAPIConfigFromFlagsMockRecorder) APIAuthKeyFromFlags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIAuthKeyFromFlags", reflect.TypeOf((*MockAPIConfigFromFlags)(nil).APIAuthKeyFromFlags))
}