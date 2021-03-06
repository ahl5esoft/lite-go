// Code generated by MockGen. DO NOT EDIT.
// Source: thread\i-lock.go

// Package thread is a generated GoMock package.
package thread

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockILock is a mock of ILock interface
type MockILock struct {
	ctrl     *gomock.Controller
	recorder *MockILockMockRecorder
}

// MockILockMockRecorder is the mock recorder for MockILock
type MockILockMockRecorder struct {
	mock *MockILock
}

// NewMockILock creates a new mock instance
func NewMockILock(ctrl *gomock.Controller) *MockILock {
	mock := &MockILock{ctrl: ctrl}
	mock.recorder = &MockILockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockILock) EXPECT() *MockILockMockRecorder {
	return m.recorder
}

// Lock mocks base method
func (m *MockILock) Lock(key string, options ...LockOption) (func(), error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Lock", varargs...)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock
func (mr *MockILockMockRecorder) Lock(key interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockILock)(nil).Lock), varargs...)
}
