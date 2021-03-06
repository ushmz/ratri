// Code generated by MockGen. DO NOT EDIT.
// Source: serp.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "ratri/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSerpUsecase is a mock of SerpUsecase interface.
type MockSerpUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockSerpUsecaseMockRecorder
}

// MockSerpUsecaseMockRecorder is the mock recorder for MockSerpUsecase.
type MockSerpUsecaseMockRecorder struct {
	mock *MockSerpUsecase
}

// NewMockSerpUsecase creates a new mock instance.
func NewMockSerpUsecase(ctrl *gomock.Controller) *MockSerpUsecase {
	mock := &MockSerpUsecase{ctrl: ctrl}
	mock.recorder = &MockSerpUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSerpUsecase) EXPECT() *MockSerpUsecaseMockRecorder {
	return m.recorder
}

// FetchSerp mocks base method.
func (m *MockSerpUsecase) FetchSerp(taskID, offset int) ([]model.SearchPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerp", taskID, offset)
	ret0, _ := ret[0].([]model.SearchPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerp indicates an expected call of FetchSerp.
func (mr *MockSerpUsecaseMockRecorder) FetchSerp(taskID, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerp", reflect.TypeOf((*MockSerpUsecase)(nil).FetchSerp), taskID, offset)
}

// FetchSerpWithIcon mocks base method.
func (m *MockSerpUsecase) FetchSerpWithIcon(taskID, offset, top int) ([]model.SerpWithIcon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerpWithIcon", taskID, offset, top)
	ret0, _ := ret[0].([]model.SerpWithIcon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerpWithIcon indicates an expected call of FetchSerpWithIcon.
func (mr *MockSerpUsecaseMockRecorder) FetchSerpWithIcon(taskID, offset, top interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerpWithIcon", reflect.TypeOf((*MockSerpUsecase)(nil).FetchSerpWithIcon), taskID, offset, top)
}

// FetchSerpWithRatio mocks base method.
func (m *MockSerpUsecase) FetchSerpWithRatio(taskID, offset, top int) ([]model.SerpWithRatio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerpWithRatio", taskID, offset, top)
	ret0, _ := ret[0].([]model.SerpWithRatio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerpWithRatio indicates an expected call of FetchSerpWithRatio.
func (mr *MockSerpUsecaseMockRecorder) FetchSerpWithRatio(taskID, offset, top interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerpWithRatio", reflect.TypeOf((*MockSerpUsecase)(nil).FetchSerpWithRatio), taskID, offset, top)
}
