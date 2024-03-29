// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-api.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIApi is a mock of IApi interface.
type MockIApi struct {
	ctrl     *gomock.Controller
	recorder *MockIApiMockRecorder
}

// MockIApiMockRecorder is the mock recorder for MockIApi.
type MockIApiMockRecorder struct {
	mock *MockIApi
}

// NewMockIApi creates a new mock instance.
func NewMockIApi(ctrl *gomock.Controller) *MockIApi {
	mock := &MockIApi{ctrl: ctrl}
	mock.recorder = &MockIApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIApi) EXPECT() *MockIApiMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockIApi) Call() (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call")
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockIApiMockRecorder) Call() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockIApi)(nil).Call))
}
