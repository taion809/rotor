// Code generated by MockGen. DO NOT EDIT.
// Source: diagnostics.go

package executor

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockDiagnosticsCallback is a mock of DiagnosticsCallback interface
type MockDiagnosticsCallback struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosticsCallbackMockRecorder
}

// MockDiagnosticsCallbackMockRecorder is the mock recorder for MockDiagnosticsCallback
type MockDiagnosticsCallbackMockRecorder struct {
	mock *MockDiagnosticsCallback
}

// NewMockDiagnosticsCallback creates a new mock instance
func NewMockDiagnosticsCallback(ctrl *gomock.Controller) *MockDiagnosticsCallback {
	mock := &MockDiagnosticsCallback{ctrl: ctrl}
	mock.recorder = &MockDiagnosticsCallbackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiagnosticsCallback) EXPECT() *MockDiagnosticsCallbackMockRecorder {
	return m.recorder
}

// TaskStarted mocks base method
func (m *MockDiagnosticsCallback) TaskStarted(arg0 int) {
	m.ctrl.Call(m, "TaskStarted", arg0)
}

// TaskStarted indicates an expected call of TaskStarted
func (mr *MockDiagnosticsCallbackMockRecorder) TaskStarted(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskStarted", reflect.TypeOf((*MockDiagnosticsCallback)(nil).TaskStarted), arg0)
}

// TaskCompleted mocks base method
func (m *MockDiagnosticsCallback) TaskCompleted(arg0 AttemptResult, arg1 time.Duration) {
	m.ctrl.Call(m, "TaskCompleted", arg0, arg1)
}

// TaskCompleted indicates an expected call of TaskCompleted
func (mr *MockDiagnosticsCallbackMockRecorder) TaskCompleted(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskCompleted", reflect.TypeOf((*MockDiagnosticsCallback)(nil).TaskCompleted), arg0, arg1)
}

// AttemptStarted mocks base method
func (m *MockDiagnosticsCallback) AttemptStarted(arg0 time.Duration) {
	m.ctrl.Call(m, "AttemptStarted", arg0)
}

// AttemptStarted indicates an expected call of AttemptStarted
func (mr *MockDiagnosticsCallbackMockRecorder) AttemptStarted(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttemptStarted", reflect.TypeOf((*MockDiagnosticsCallback)(nil).AttemptStarted), arg0)
}

// AttemptCompleted mocks base method
func (m *MockDiagnosticsCallback) AttemptCompleted(arg0 AttemptResult, arg1 time.Duration) {
	m.ctrl.Call(m, "AttemptCompleted", arg0, arg1)
}

// AttemptCompleted indicates an expected call of AttemptCompleted
func (mr *MockDiagnosticsCallbackMockRecorder) AttemptCompleted(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttemptCompleted", reflect.TypeOf((*MockDiagnosticsCallback)(nil).AttemptCompleted), arg0, arg1)
}

// CallbackDuration mocks base method
func (m *MockDiagnosticsCallback) CallbackDuration(arg0 time.Duration) {
	m.ctrl.Call(m, "CallbackDuration", arg0)
}

// CallbackDuration indicates an expected call of CallbackDuration
func (mr *MockDiagnosticsCallbackMockRecorder) CallbackDuration(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallbackDuration", reflect.TypeOf((*MockDiagnosticsCallback)(nil).CallbackDuration), arg0)
}