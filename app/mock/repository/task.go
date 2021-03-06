// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "ratri/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// CreateTaskAnswer mocks base method.
func (m *MockTaskRepository) CreateTaskAnswer(arg0 *model.Answer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTaskAnswer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTaskAnswer indicates an expected call of CreateTaskAnswer.
func (mr *MockTaskRepositoryMockRecorder) CreateTaskAnswer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTaskAnswer", reflect.TypeOf((*MockTaskRepository)(nil).CreateTaskAnswer), arg0)
}

// FetchTaskInfo mocks base method.
func (m *MockTaskRepository) FetchTaskInfo(taskID int) (*model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTaskInfo", taskID)
	ret0, _ := ret[0].(*model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTaskInfo indicates an expected call of FetchTaskInfo.
func (mr *MockTaskRepositoryMockRecorder) FetchTaskInfo(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTaskInfo", reflect.TypeOf((*MockTaskRepository)(nil).FetchTaskInfo), taskID)
}

// GetConditionIDByGroupID mocks base method.
func (m *MockTaskRepository) GetConditionIDByGroupID(groupID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConditionIDByGroupID", groupID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConditionIDByGroupID indicates an expected call of GetConditionIDByGroupID.
func (mr *MockTaskRepositoryMockRecorder) GetConditionIDByGroupID(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConditionIDByGroupID", reflect.TypeOf((*MockTaskRepository)(nil).GetConditionIDByGroupID), groupID)
}

// GetTaskIDsByGroupID mocks base method.
func (m *MockTaskRepository) GetTaskIDsByGroupID(groupID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskIDsByGroupID", groupID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskIDsByGroupID indicates an expected call of GetTaskIDsByGroupID.
func (mr *MockTaskRepositoryMockRecorder) GetTaskIDsByGroupID(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskIDsByGroupID", reflect.TypeOf((*MockTaskRepository)(nil).GetTaskIDsByGroupID), groupID)
}

// UpdateTaskAllocation mocks base method.
func (m *MockTaskRepository) UpdateTaskAllocation() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskAllocation")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskAllocation indicates an expected call of UpdateTaskAllocation.
func (mr *MockTaskRepositoryMockRecorder) UpdateTaskAllocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskAllocation", reflect.TypeOf((*MockTaskRepository)(nil).UpdateTaskAllocation))
}
