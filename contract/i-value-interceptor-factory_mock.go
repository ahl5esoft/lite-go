// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-value-interceptor-factory.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	contract "github.com/ahl5esoft/lite-go/model/contract"
	message "github.com/ahl5esoft/lite-go/model/message"
	gomock "github.com/golang/mock/gomock"
)

// MockIValueInterceptorFactory is a mock of IValueInterceptorFactory interface.
type MockIValueInterceptorFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIValueInterceptorFactoryMockRecorder
}

// MockIValueInterceptorFactoryMockRecorder is the mock recorder for MockIValueInterceptorFactory.
type MockIValueInterceptorFactoryMockRecorder struct {
	mock *MockIValueInterceptorFactory
}

// NewMockIValueInterceptorFactory creates a new mock instance.
func NewMockIValueInterceptorFactory(ctrl *gomock.Controller) *MockIValueInterceptorFactory {
	mock := &MockIValueInterceptorFactory{ctrl: ctrl}
	mock.recorder = &MockIValueInterceptorFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIValueInterceptorFactory) EXPECT() *MockIValueInterceptorFactoryMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockIValueInterceptorFactory) Build(arg0 message.ChangeValue) (IValueInterceptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(IValueInterceptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockIValueInterceptorFactoryMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockIValueInterceptorFactory)(nil).Build), arg0)
}

// Register mocks base method.
func (m *MockIValueInterceptorFactory) Register(arg0 int, arg1 IValueInterceptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0, arg1)
}

// Register indicates an expected call of Register.
func (mr *MockIValueInterceptorFactoryMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIValueInterceptorFactory)(nil).Register), arg0, arg1)
}

// RegisterPredicate mocks base method.
func (m *MockIValueInterceptorFactory) RegisterPredicate(arg0 func(contract.IEnumItem) bool, arg1 IValueInterceptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterPredicate", arg0, arg1)
}

// RegisterPredicate indicates an expected call of RegisterPredicate.
func (mr *MockIValueInterceptorFactoryMockRecorder) RegisterPredicate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPredicate", reflect.TypeOf((*MockIValueInterceptorFactory)(nil).RegisterPredicate), arg0, arg1)
}
