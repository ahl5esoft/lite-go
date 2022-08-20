// Code generated by MockGen. DO NOT EDIT.
// Source: contract/i-os-path.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIOsPath is a mock of IOsPath interface.
type MockIOsPath struct {
	ctrl     *gomock.Controller
	recorder *MockIOsPathMockRecorder
}

// MockIOsPathMockRecorder is the mock recorder for MockIOsPath.
type MockIOsPathMockRecorder struct {
	mock *MockIOsPath
}

// NewMockIOsPath creates a new mock instance.
func NewMockIOsPath(ctrl *gomock.Controller) *MockIOsPath {
	mock := &MockIOsPath{ctrl: ctrl}
	mock.recorder = &MockIOsPathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOsPath) EXPECT() *MockIOsPathMockRecorder {
	return m.recorder
}

// GetRoot mocks base method.
func (m *MockIOsPath) GetRoot() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoot")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRoot indicates an expected call of GetRoot.
func (mr *MockIOsPathMockRecorder) GetRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoot", reflect.TypeOf((*MockIOsPath)(nil).GetRoot))
}

// Join mocks base method.
func (m *MockIOsPath) Join(paths ...string) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range paths {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Join", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Join indicates an expected call of Join.
func (mr *MockIOsPathMockRecorder) Join(paths ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockIOsPath)(nil).Join), paths...)
}
