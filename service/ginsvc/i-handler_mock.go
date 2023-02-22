// Code generated by MockGen. DO NOT EDIT.
// Source: service\ginsvc\i-handler.go

// Package ginsvc is a generated GoMock package.
package ginsvc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIHandler is a mock of IHandler interface.
type MockIHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIHandlerMockRecorder
}

// MockIHandlerMockRecorder is the mock recorder for MockIHandler.
type MockIHandlerMockRecorder struct {
	mock *MockIHandler
}

// NewMockIHandler creates a new mock instance.
func NewMockIHandler(ctrl *gomock.Controller) *MockIHandler {
	mock := &MockIHandler{ctrl: ctrl}
	mock.recorder = &MockIHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHandler) EXPECT() *MockIHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockIHandler) Handle(v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockIHandlerMockRecorder) Handle(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIHandler)(nil).Handle), v)
}

// SetNext mocks base method.
func (m *MockIHandler) SetNext(next IHandler) IHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNext", next)
	ret0, _ := ret[0].(IHandler)
	return ret0
}

// SetNext indicates an expected call of SetNext.
func (mr *MockIHandlerMockRecorder) SetNext(next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNext", reflect.TypeOf((*MockIHandler)(nil).SetNext), next)
}