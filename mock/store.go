// Code generated by MockGen. DO NOT EDIT.
// Source: cthulhu/store (interfaces: Service)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	telegram "cthulhu/telegram"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// StoreService is a mock of Service interface
type StoreService struct {
	ctrl     *gomock.Controller
	recorder *StoreServiceMockRecorder
}

// StoreServiceMockRecorder is the mock recorder for StoreService
type StoreServiceMockRecorder struct {
	mock *StoreService
}

// NewStoreService creates a new mock instance
func NewStoreService(ctrl *gomock.Controller) *StoreService {
	mock := &StoreService{ctrl: ctrl}
	mock.recorder = &StoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *StoreService) EXPECT() *StoreServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *StoreService) Create(arg0 context.Context, arg1 string, arg2 *telegram.Update) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *StoreServiceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*StoreService)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *StoreService) Delete(arg0 context.Context, arg1 string) (*telegram.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*telegram.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *StoreServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*StoreService)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method
func (m *StoreService) GetAll(arg0 context.Context) map[string]*telegram.Update {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(map[string]*telegram.Update)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *StoreServiceMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*StoreService)(nil).GetAll), arg0)
}

// Read mocks base method
func (m *StoreService) Read(arg0 context.Context, arg1 string) (*telegram.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*telegram.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *StoreServiceMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*StoreService)(nil).Read), arg0, arg1)
}

// Update mocks base method
func (m *StoreService) Update(arg0 context.Context, arg1 string, arg2 *telegram.Update) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *StoreServiceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*StoreService)(nil).Update), arg0, arg1, arg2)
}
