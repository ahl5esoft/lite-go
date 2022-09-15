// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-db-factory.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	contract "github.com/ahl5esoft/lite-go/model/contract"
	gomock "github.com/golang/mock/gomock"
)

// MockIDbFactory is a mock of IDbFactory interface.
type MockIDbFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIDbFactoryMockRecorder
}

// MockIDbFactoryMockRecorder is the mock recorder for MockIDbFactory.
type MockIDbFactoryMockRecorder struct {
	mock *MockIDbFactory
}

// NewMockIDbFactory creates a new mock instance.
func NewMockIDbFactory(ctrl *gomock.Controller) *MockIDbFactory {
	mock := &MockIDbFactory{ctrl: ctrl}
	mock.recorder = &MockIDbFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDbFactory) EXPECT() *MockIDbFactoryMockRecorder {
	return m.recorder
}

// Db mocks base method.
func (m *MockIDbFactory) Db(entry contract.IDbModel, extra ...any) IDbRepository {
	m.ctrl.T.Helper()
	varargs := []interface{}{entry}
	for _, a := range extra {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Db", varargs...)
	ret0, _ := ret[0].(IDbRepository)
	return ret0
}

// Db indicates an expected call of Db.
func (mr *MockIDbFactoryMockRecorder) Db(entry interface{}, extra ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{entry}, extra...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Db", reflect.TypeOf((*MockIDbFactory)(nil).Db), varargs...)
}

// Uow mocks base method.
func (m *MockIDbFactory) Uow() IUnitOfWork {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uow")
	ret0, _ := ret[0].(IUnitOfWork)
	return ret0
}

// Uow indicates an expected call of Uow.
func (mr *MockIDbFactoryMockRecorder) Uow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uow", reflect.TypeOf((*MockIDbFactory)(nil).Uow))
}
