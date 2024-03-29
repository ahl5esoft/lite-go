// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-user-reward-service.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	message "github.com/ahl5esoft/lite-go/model/message"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserRewardService is a mock of IUserRewardService interface.
type MockIUserRewardService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRewardServiceMockRecorder
}

// MockIUserRewardServiceMockRecorder is the mock recorder for MockIUserRewardService.
type MockIUserRewardServiceMockRecorder struct {
	mock *MockIUserRewardService
}

// NewMockIUserRewardService creates a new mock instance.
func NewMockIUserRewardService(ctrl *gomock.Controller) *MockIUserRewardService {
	mock := &MockIUserRewardService{ctrl: ctrl}
	mock.recorder = &MockIUserRewardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRewardService) EXPECT() *MockIUserRewardServiceMockRecorder {
	return m.recorder
}

// FindResults mocks base method.
func (m *MockIUserRewardService) FindResults(arg0 IUnitOfWork, arg1 [][]message.Reward, arg2, arg3 string) ([]message.ChangeValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindResults", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]message.ChangeValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResults indicates an expected call of FindResults.
func (mr *MockIUserRewardServiceMockRecorder) FindResults(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResults", reflect.TypeOf((*MockIUserRewardService)(nil).FindResults), arg0, arg1, arg2, arg3)
}

// Preview mocks base method.
func (m *MockIUserRewardService) Preview(arg0 IUnitOfWork, arg1 [][]message.Reward, arg2 string) ([]message.ChangeValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Preview", arg0, arg1, arg2)
	ret0, _ := ret[0].([]message.ChangeValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Preview indicates an expected call of Preview.
func (mr *MockIUserRewardServiceMockRecorder) Preview(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preview", reflect.TypeOf((*MockIUserRewardService)(nil).Preview), arg0, arg1, arg2)
}
