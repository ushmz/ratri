// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "ratri/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// AllocateTask mocks base method.
func (m *MockUserUsecase) AllocateTask() (model.TaskInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateTask")
	ret0, _ := ret[0].(model.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocateTask indicates an expected call of AllocateTask.
func (mr *MockUserUsecaseMockRecorder) AllocateTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateTask", reflect.TypeOf((*MockUserUsecase)(nil).AllocateTask))
}

// CreateUser mocks base method.
func (m *MockUserUsecase) CreateUser(uid string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", uid)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUsecaseMockRecorder) CreateUser(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUsecase)(nil).CreateUser), uid)
}

// FindByUid mocks base method.
func (m *MockUserUsecase) FindByUid(uid string) (model.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUid", uid)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindByUid indicates an expected call of FindByUid.
func (mr *MockUserUsecaseMockRecorder) FindByUid(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUid", reflect.TypeOf((*MockUserUsecase)(nil).FindByUid), uid)
}

// GenerateRandomPasswd mocks base method.
func (m *MockUserUsecase) GenerateRandomPasswd(l int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomPasswd", l)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateRandomPasswd indicates an expected call of GenerateRandomPasswd.
func (mr *MockUserUsecaseMockRecorder) GenerateRandomPasswd(l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomPasswd", reflect.TypeOf((*MockUserUsecase)(nil).GenerateRandomPasswd), l)
}

// GetCompletionCode mocks base method.
func (m *MockUserUsecase) GetCompletionCode(userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompletionCode", userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompletionCode indicates an expected call of GetCompletionCode.
func (mr *MockUserUsecaseMockRecorder) GetCompletionCode(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletionCode", reflect.TypeOf((*MockUserUsecase)(nil).GetCompletionCode), userId)
}
