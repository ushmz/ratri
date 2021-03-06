// Code generated by MockGen. DO NOT EDIT.
// Source: serp.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "ratri/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSerpRepository is a mock of SerpRepository interface.
type MockSerpRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSerpRepositoryMockRecorder
}

// MockSerpRepositoryMockRecorder is the mock recorder for MockSerpRepository.
type MockSerpRepositoryMockRecorder struct {
	mock *MockSerpRepository
}

// NewMockSerpRepository creates a new mock instance.
func NewMockSerpRepository(ctrl *gomock.Controller) *MockSerpRepository {
	mock := &MockSerpRepository{ctrl: ctrl}
	mock.recorder = &MockSerpRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSerpRepository) EXPECT() *MockSerpRepositoryMockRecorder {
	return m.recorder
}

// FetchSerpByTaskID mocks base method.
func (m *MockSerpRepository) FetchSerpByTaskID(taskID, offset int) ([]model.SearchPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerpByTaskID", taskID, offset)
	ret0, _ := ret[0].([]model.SearchPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerpByTaskID indicates an expected call of FetchSerpByTaskID.
func (mr *MockSerpRepositoryMockRecorder) FetchSerpByTaskID(taskID, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerpByTaskID", reflect.TypeOf((*MockSerpRepository)(nil).FetchSerpByTaskID), taskID, offset)
}
