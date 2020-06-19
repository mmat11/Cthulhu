// Code generated by MockGen. DO NOT EDIT.
// Source: cthulhu/telegram (interfaces: Service)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// TelegramService is a mock of Service interface
type TelegramService struct {
	ctrl     *gomock.Controller
	recorder *TelegramServiceMockRecorder
}

// TelegramServiceMockRecorder is the mock recorder for TelegramService
type TelegramServiceMockRecorder struct {
	mock *TelegramService
}

// NewTelegramService creates a new mock instance
func NewTelegramService(ctrl *gomock.Controller) *TelegramService {
	mock := &TelegramService{ctrl: ctrl}
	mock.recorder = &TelegramServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *TelegramService) EXPECT() *TelegramServiceMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method
func (m *TelegramService) DeleteMessage(arg0 context.Context, arg1 int64, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage
func (mr *TelegramServiceMockRecorder) DeleteMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*TelegramService)(nil).DeleteMessage), arg0, arg1, arg2)
}

// KickChatMember mocks base method
func (m *TelegramService) KickChatMember(arg0 context.Context, arg1 int64, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KickChatMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// KickChatMember indicates an expected call of KickChatMember
func (mr *TelegramServiceMockRecorder) KickChatMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KickChatMember", reflect.TypeOf((*TelegramService)(nil).KickChatMember), arg0, arg1, arg2)
}

// Reply mocks base method
func (m *TelegramService) Reply(arg0 context.Context, arg1 int64, arg2 string, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reply indicates an expected call of Reply
func (mr *TelegramServiceMockRecorder) Reply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reply", reflect.TypeOf((*TelegramService)(nil).Reply), arg0, arg1, arg2, arg3)
}

// SendMessage mocks base method
func (m *TelegramService) SendMessage(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (mr *TelegramServiceMockRecorder) SendMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*TelegramService)(nil).SendMessage), arg0, arg1, arg2)
}

// UnbanChatMember mocks base method
func (m *TelegramService) UnbanChatMember(arg0 context.Context, arg1 int64, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbanChatMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnbanChatMember indicates an expected call of UnbanChatMember
func (mr *TelegramServiceMockRecorder) UnbanChatMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbanChatMember", reflect.TypeOf((*TelegramService)(nil).UnbanChatMember), arg0, arg1, arg2)
}
