// Code generated by MockGen. DO NOT EDIT.
// Source: cryptoex\i-crypto.go

// Package cryptoex is a generated GoMock package.
package cryptoex

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICrypto is a mock of ICrypto interface
type MockICrypto struct {
	ctrl     *gomock.Controller
	recorder *MockICryptoMockRecorder
}

// MockICryptoMockRecorder is the mock recorder for MockICrypto
type MockICryptoMockRecorder struct {
	mock *MockICrypto
}

// NewMockICrypto creates a new mock instance
func NewMockICrypto(ctrl *gomock.Controller) *MockICrypto {
	mock := &MockICrypto{ctrl: ctrl}
	mock.recorder = &MockICryptoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICrypto) EXPECT() *MockICryptoMockRecorder {
	return m.recorder
}

// Decrypt mocks base method
func (m *MockICrypto) Decrypt(ciphertext []byte, options ...DecryptOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ciphertext}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Decrypt", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt
func (mr *MockICryptoMockRecorder) Decrypt(ciphertext interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ciphertext}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockICrypto)(nil).Decrypt), varargs...)
}

// Encrypt mocks base method
func (m *MockICrypto) Encrypt(plaintext []byte, options ...EncryptOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{plaintext}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Encrypt", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockICryptoMockRecorder) Encrypt(plaintext interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{plaintext}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockICrypto)(nil).Encrypt), varargs...)
}

// Validate mocks base method
func (m *MockICrypto) Validate(ciphertext, plaintext []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ciphertext, plaintext)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockICryptoMockRecorder) Validate(ciphertext, plaintext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockICrypto)(nil).Validate), ciphertext, plaintext)
}
