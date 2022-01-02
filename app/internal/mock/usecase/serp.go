// Code generated by MockGen. DO NOT EDIT.
// Source: serp.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "ratri/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSerp is a mock of Serp interface.
type MockSerp struct {
	ctrl     *gomock.Controller
	recorder *MockSerpMockRecorder
}

// MockSerpMockRecorder is the mock recorder for MockSerp.
type MockSerpMockRecorder struct {
	mock *MockSerp
}

// NewMockSerp creates a new mock instance.
func NewMockSerp(ctrl *gomock.Controller) *MockSerp {
	mock := &MockSerp{ctrl: ctrl}
	mock.recorder = &MockSerpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSerp) EXPECT() *MockSerpMockRecorder {
	return m.recorder
}

// FetchSerp mocks base method.
func (m *MockSerp) FetchSerp(taskId, offset int) (*[]model.SearchPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerp", taskId, offset)
	ret0, _ := ret[0].(*[]model.SearchPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerp indicates an expected call of FetchSerp.
func (mr *MockSerpMockRecorder) FetchSerp(taskId, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerp", reflect.TypeOf((*MockSerp)(nil).FetchSerp), taskId, offset)
}

// FetchSerpWithIcon mocks base method.
func (m *MockSerp) FetchSerpWithIcon(taskId, offset, top int) (*[]model.SerpWithIcon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerpWithIcon", taskId, offset, top)
	ret0, _ := ret[0].(*[]model.SerpWithIcon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerpWithIcon indicates an expected call of FetchSerpWithIcon.
func (mr *MockSerpMockRecorder) FetchSerpWithIcon(taskId, offset, top interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerpWithIcon", reflect.TypeOf((*MockSerp)(nil).FetchSerpWithIcon), taskId, offset, top)
}

// FetchSerpWithRatio mocks base method.
func (m *MockSerp) FetchSerpWithRatio(taskId, offset, top int) (*[]model.SerpWithRatio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSerpWithRatio", taskId, offset, top)
	ret0, _ := ret[0].(*[]model.SerpWithRatio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSerpWithRatio indicates an expected call of FetchSerpWithRatio.
func (mr *MockSerpMockRecorder) FetchSerpWithRatio(taskId, offset, top interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSerpWithRatio", reflect.TypeOf((*MockSerp)(nil).FetchSerpWithRatio), taskId, offset, top)
}
