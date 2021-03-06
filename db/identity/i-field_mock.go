// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ahl5esoft/lite-go/db/identity (interfaces: IField)

// Package identity is a generated GoMock package.
package identity

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIField is a mock of IField interface
type MockIField struct {
	ctrl     *gomock.Controller
	recorder *MockIFieldMockRecorder
}

// MockIFieldMockRecorder is the mock recorder for MockIField
type MockIFieldMockRecorder struct {
	mock *MockIField
}

// NewMockIField creates a new mock instance
func NewMockIField(ctrl *gomock.Controller) *MockIField {
	mock := &MockIField{ctrl: ctrl}
	mock.recorder = &MockIFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIField) EXPECT() *MockIFieldMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockIField) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockIFieldMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIField)(nil).GetName))
}

// GetStructName mocks base method
func (m *MockIField) GetStructName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStructName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStructName indicates an expected call of GetStructName
func (mr *MockIFieldMockRecorder) GetStructName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStructName", reflect.TypeOf((*MockIField)(nil).GetStructName))
}

// GetValue mocks base method
func (m *MockIField) GetValue(arg0 reflect.Value) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetValue indicates an expected call of GetValue
func (mr *MockIFieldMockRecorder) GetValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockIField)(nil).GetValue), arg0)
}
