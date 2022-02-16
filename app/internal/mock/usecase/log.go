// Code generated by MockGen. DO NOT EDIT.
// Source: log.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "ratri/internal/domain/model"
	store "ratri/internal/domain/store"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogUsecase is a mock of LogUsecase interface.
type MockLogUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockLogUsecaseMockRecorder
}

// MockLogUsecaseMockRecorder is the mock recorder for MockLogUsecase.
type MockLogUsecaseMockRecorder struct {
	mock *MockLogUsecase
}

// NewMockLogUsecase creates a new mock instance.
func NewMockLogUsecase(ctrl *gomock.Controller) *MockLogUsecase {
	mock := &MockLogUsecase{ctrl: ctrl}
	mock.recorder = &MockLogUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogUsecase) EXPECT() *MockLogUsecaseMockRecorder {
	return m.recorder
}

// CumulatePageViewingTime mocks base method.
func (m *MockLogUsecase) CumulatePageViewingTime(arg0 *model.PageViewingLogParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CumulatePageViewingTime", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CumulatePageViewingTime indicates an expected call of CumulatePageViewingTime.
func (mr *MockLogUsecaseMockRecorder) CumulatePageViewingTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CumulatePageViewingTime", reflect.TypeOf((*MockLogUsecase)(nil).CumulatePageViewingTime), arg0)
}

// CumulateSerpViewingTime mocks base method.
func (m *MockLogUsecase) CumulateSerpViewingTime(arg0 *model.SerpViewingLogParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CumulateSerpViewingTime", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CumulateSerpViewingTime indicates an expected call of CumulateSerpViewingTime.
func (mr *MockLogUsecaseMockRecorder) CumulateSerpViewingTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CumulateSerpViewingTime", reflect.TypeOf((*MockLogUsecase)(nil).CumulateSerpViewingTime), arg0)
}

// ExportPageViewingTimeLog mocks base method.
func (m *MockLogUsecase) ExportPageViewingTimeLog(arg0 bool, arg1 store.FileType) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportPageViewingTimeLog", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportPageViewingTimeLog indicates an expected call of ExportPageViewingTimeLog.
func (mr *MockLogUsecaseMockRecorder) ExportPageViewingTimeLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportPageViewingTimeLog", reflect.TypeOf((*MockLogUsecase)(nil).ExportPageViewingTimeLog), arg0, arg1)
}

// ExportSearchSeeion mocks base method.
func (m *MockLogUsecase) ExportSearchSeeion(arg0 bool, arg1 store.FileType) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSearchSeeion", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSearchSeeion indicates an expected call of ExportSearchSeeion.
func (mr *MockLogUsecaseMockRecorder) ExportSearchSeeion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSearchSeeion", reflect.TypeOf((*MockLogUsecase)(nil).ExportSearchSeeion), arg0, arg1)
}

// ExportSerpEventLog mocks base method.
func (m *MockLogUsecase) ExportSerpEventLog(arg0 bool, arg1 store.FileType) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSerpEventLog", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSerpEventLog indicates an expected call of ExportSerpEventLog.
func (mr *MockLogUsecaseMockRecorder) ExportSerpEventLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSerpEventLog", reflect.TypeOf((*MockLogUsecase)(nil).ExportSerpEventLog), arg0, arg1)
}

// ExportSerpViewingTimeLog mocks base method.
func (m *MockLogUsecase) ExportSerpViewingTimeLog(arg0 bool, arg1 store.FileType) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSerpViewingTimeLog", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSerpViewingTimeLog indicates an expected call of ExportSerpViewingTimeLog.
func (mr *MockLogUsecaseMockRecorder) ExportSerpViewingTimeLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSerpViewingTimeLog", reflect.TypeOf((*MockLogUsecase)(nil).ExportSerpViewingTimeLog), arg0, arg1)
}

// StoreSearchSeeion mocks base method.
func (m *MockLogUsecase) StoreSearchSeeion(arg0 *model.SearchSessionParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSearchSeeion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSearchSeeion indicates an expected call of StoreSearchSeeion.
func (mr *MockLogUsecaseMockRecorder) StoreSearchSeeion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSearchSeeion", reflect.TypeOf((*MockLogUsecase)(nil).StoreSearchSeeion), arg0)
}

// StoreSerpEventLog mocks base method.
func (m *MockLogUsecase) StoreSerpEventLog(arg0 *model.SearchPageEventLogParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSerpEventLog", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSerpEventLog indicates an expected call of StoreSerpEventLog.
func (mr *MockLogUsecaseMockRecorder) StoreSerpEventLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSerpEventLog", reflect.TypeOf((*MockLogUsecase)(nil).StoreSerpEventLog), arg0)
}
