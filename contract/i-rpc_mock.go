// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-rpc.go

// Package contract is a generated GoMock package.
package contract

import (
	message "github.com/ahl5esoft/lite-go/model/message"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRpc is a mock of IRpc interface
type MockIRpc struct {
	ctrl     *gomock.Controller
	recorder *MockIRpcMockRecorder
}

// MockIRpcMockRecorder is the mock recorder for MockIRpc
type MockIRpcMockRecorder struct {
	mock *MockIRpc
}

// NewMockIRpc creates a new mock instance
func NewMockIRpc(ctrl *gomock.Controller) *MockIRpc {
	mock := &MockIRpc{ctrl: ctrl}
	mock.recorder = &MockIRpcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRpc) EXPECT() *MockIRpcMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockIRpc) Call(arg0 string, arg1 *message.ApiResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Call indicates an expected call of Call
func (mr *MockIRpcMockRecorder) Call(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockIRpc)(nil).Call), arg0, arg1)
}

// SetBody mocks base method
func (m *MockIRpc) SetBody(arg0 any) IRpc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBody", arg0)
	ret0, _ := ret[0].(IRpc)
	return ret0
}

// SetBody indicates an expected call of SetBody
func (mr *MockIRpcMockRecorder) SetBody(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBody", reflect.TypeOf((*MockIRpc)(nil).SetBody), arg0)
}

// SetHeader mocks base method
func (m *MockIRpc) SetHeader(arg0 map[string]string) IRpc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(IRpc)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockIRpcMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockIRpc)(nil).SetHeader), arg0)
}
