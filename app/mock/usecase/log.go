// Code generated by MockGen. DO NOT EDIT.
// Source: log.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	bytes "bytes"
	model "ratri/domain/model"
	store "ratri/domain/store"
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

// CumulatePageDwellTime mocks base method.
func (m *MockLogUsecase) CumulatePageDwellTime(arg0 *model.PageDwellTimeLogParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CumulatePageDwellTime", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CumulatePageDwellTime indicates an expected call of CumulatePageDwellTime.
func (mr *MockLogUsecaseMockRecorder) CumulatePageDwellTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CumulatePageDwellTime", reflect.TypeOf((*MockLogUsecase)(nil).CumulatePageDwellTime), arg0)
}

// CumulateSerpDwellTime mocks base method.
func (m *MockLogUsecase) CumulateSerpDwellTime(arg0 *model.SerpDwellTimeLogParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CumulateSerpDwellTime", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CumulateSerpDwellTime indicates an expected call of CumulateSerpDwellTime.
func (mr *MockLogUsecaseMockRecorder) CumulateSerpDwellTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CumulateSerpDwellTime", reflect.TypeOf((*MockLogUsecase)(nil).CumulateSerpDwellTime), arg0)
}

// ExportPageDwellTimeLog mocks base method.
func (m *MockLogUsecase) ExportPageDwellTimeLog(arg0 bool, arg1 store.FileType) (*bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportPageDwellTimeLog", arg0, arg1)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportPageDwellTimeLog indicates an expected call of ExportPageDwellTimeLog.
func (mr *MockLogUsecaseMockRecorder) ExportPageDwellTimeLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportPageDwellTimeLog", reflect.TypeOf((*MockLogUsecase)(nil).ExportPageDwellTimeLog), arg0, arg1)
}

// ExportSearchSeeion mocks base method.
func (m *MockLogUsecase) ExportSearchSeeion(arg0 bool, arg1 store.FileType) (*bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSearchSeeion", arg0, arg1)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSearchSeeion indicates an expected call of ExportSearchSeeion.
func (mr *MockLogUsecaseMockRecorder) ExportSearchSeeion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSearchSeeion", reflect.TypeOf((*MockLogUsecase)(nil).ExportSearchSeeion), arg0, arg1)
}

// ExportSerpDwellTimeLog mocks base method.
func (m *MockLogUsecase) ExportSerpDwellTimeLog(arg0 bool, arg1 store.FileType) (*bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSerpDwellTimeLog", arg0, arg1)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSerpDwellTimeLog indicates an expected call of ExportSerpDwellTimeLog.
func (mr *MockLogUsecaseMockRecorder) ExportSerpDwellTimeLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSerpDwellTimeLog", reflect.TypeOf((*MockLogUsecase)(nil).ExportSerpDwellTimeLog), arg0, arg1)
}

// ExportSerpEventLog mocks base method.
func (m *MockLogUsecase) ExportSerpEventLog(arg0 bool, arg1 store.FileType) (*bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSerpEventLog", arg0, arg1)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSerpEventLog indicates an expected call of ExportSerpEventLog.
func (mr *MockLogUsecaseMockRecorder) ExportSerpEventLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSerpEventLog", reflect.TypeOf((*MockLogUsecase)(nil).ExportSerpEventLog), arg0, arg1)
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
