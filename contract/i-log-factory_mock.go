// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-log-factory.go

// Package contract is a generated GoMock package.
package contract

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockILogFactory is a mock of ILogFactory interface
type MockILogFactory struct {
	ctrl     *gomock.Controller
	recorder *MockILogFactoryMockRecorder
}

// MockILogFactoryMockRecorder is the mock recorder for MockILogFactory
type MockILogFactoryMockRecorder struct {
	mock *MockILogFactory
}

// NewMockILogFactory creates a new mock instance
func NewMockILogFactory(ctrl *gomock.Controller) *MockILogFactory {
	mock := &MockILogFactory{ctrl: ctrl}
	mock.recorder = &MockILogFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockILogFactory) EXPECT() *MockILogFactoryMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockILogFactory) Build() ILog {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(ILog)
	return ret0
}

// Build indicates an expected call of Build
func (mr *MockILogFactoryMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockILogFactory)(nil).Build))
}
