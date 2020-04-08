// Code generated by MockGen. DO NOT EDIT.
// Source: cthulhu/bot (interfaces: Service)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	bot "cthulhu/bot"
	telegram "cthulhu/telegram"
)

// BotService is a mock of Service interface
type BotService struct {
	ctrl     *gomock.Controller
	recorder *BotServiceMockRecorder
}

// BotServiceMockRecorder is the mock recorder for BotService
type BotServiceMockRecorder struct {
	mock *BotService
}

// NewBotService creates a new mock instance
func NewBotService(ctrl *gomock.Controller) *BotService {
	mock := &BotService{ctrl: ctrl}
	mock.recorder = &BotServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *BotService) EXPECT() *BotServiceMockRecorder {
	return m.recorder
}

// GetToken mocks base method
func (m *BotService) GetToken() bot.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken")
	ret0, _ := ret[0].(bot.Token)
	return ret0
}

// GetToken indicates an expected call of GetToken
func (mr *BotServiceMockRecorder) GetToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*BotService)(nil).GetToken))
}

// Update mocks base method
func (m *BotService) Update(arg0 context.Context, arg1 *telegram.Update) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *BotServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*BotService)(nil).Update), arg0, arg1)
}
