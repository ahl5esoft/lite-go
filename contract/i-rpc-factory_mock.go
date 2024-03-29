// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-rpc-factory.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRpcFactory is a mock of IRpcFactory interface.
type MockIRpcFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIRpcFactoryMockRecorder
}

// MockIRpcFactoryMockRecorder is the mock recorder for MockIRpcFactory.
type MockIRpcFactoryMockRecorder struct {
	mock *MockIRpcFactory
}

// NewMockIRpcFactory creates a new mock instance.
func NewMockIRpcFactory(ctrl *gomock.Controller) *MockIRpcFactory {
	mock := &MockIRpcFactory{ctrl: ctrl}
	mock.recorder = &MockIRpcFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRpcFactory) EXPECT() *MockIRpcFactoryMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockIRpcFactory) Build() IRpc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(IRpc)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockIRpcFactoryMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockIRpcFactory)(nil).Build))
}
