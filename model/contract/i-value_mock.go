// Code generated by MockGen. DO NOT EDIT.
// Source: model\contract\i-value.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIValue is a mock of IValue interface.
type MockIValue struct {
	ctrl     *gomock.Controller
	recorder *MockIValueMockRecorder
}

// MockIValueMockRecorder is the mock recorder for MockIValue.
type MockIValueMockRecorder struct {
	mock *MockIValue
}

// NewMockIValue creates a new mock instance.
func NewMockIValue(ctrl *gomock.Controller) *MockIValue {
	mock := &MockIValue{ctrl: ctrl}
	mock.recorder = &MockIValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIValue) EXPECT() *MockIValueMockRecorder {
	return m.recorder
}

// GetID mocks base method.
func (m *MockIValue) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockIValueMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockIValue)(nil).GetID))
}

// GetValue mocks base method.
func (m *MockIValue) GetValue() map[int]int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue")
	ret0, _ := ret[0].(map[int]int64)
	return ret0
}

// GetValue indicates an expected call of GetValue.
func (mr *MockIValueMockRecorder) GetValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockIValue)(nil).GetValue))
}