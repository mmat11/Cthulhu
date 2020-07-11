package bot

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"cthulhu/store"
	"cthulhu/telegram"
)

const pingCommand = "ping"

type Service interface {
	Update(ctx context.Context, req *telegram.Update) error
	GetToken() Token
}

type service struct {
	Logger   log.Logger
	Config   Config
	Store    store.Service
	Telegram telegram.Service
}

func NewService(logger log.Logger, config Config, storeService store.Service, telegramService telegram.Service) *service {
	return &service{
		Logger:   logger,
		Config:   config,
		Store:    storeService,
		Telegram: telegramService,
	}
}

func (s *service) GetToken() Token {
	return s.Config.Bot.Token
}

func (s *service) Update(ctx context.Context, updateReq *telegram.Update) error {
	if updateReq.Message == nil {
		return nil
	}

	if !s.checkOrigin(ctx, updateReq) {
		level.Error(s.Logger).Log("msg", "group is not part of the network")
		return nil
	}

	s.Store.Create(ctx, strconv.Itoa(updateReq.UpdateID), telegram.MarshalUpdate(updateReq))

	if updateReq.Message.NewChatMembers != nil {
		s.handleNewUsers(ctx, updateReq)
	}

	if command := updateReq.Message.Command(); command != "" {
		level.Info(s.Logger).Log("msg", "received new command", "command", command)
		switch command {
		case pingCommand:
			return s.Telegram.SendMessage(ctx, updateReq.Message.Chat.ID, "pong")
		case banCommand:
			return s.handleBan(ctx, updateReq)
		case unbanCommand:
			return s.handleUnban(ctx, updateReq)
		case broadcastCommand:
			return s.handleBroadcast(ctx, updateReq)
		}
	}
	return s.handleCrossposts(ctx, updateReq)
}

func (s *service) checkOrigin(ctx context.Context, updateReq *telegram.Update) bool {
	var chatID int64 = updateReq.Message.Chat.ID

	for _, g := range s.Config.Bot.AccessControl.Groups {
		if g.Group.ID == chatID {
			return true
		}
	}
	return false
}

func (s *service) handleNewUsers(ctx context.Context, updateReq *telegram.Update) {
	var chatID int64 = updateReq.Message.Chat.ID

	for _, g := range s.Config.Bot.AccessControl.Groups {
		if g.Group.ID == chatID {
			if g.Group.WelcomeMessage != "" {
				for range *updateReq.Message.NewChatMembers {
					s.Telegram.Reply(ctx, g.Group.ID, g.Group.WelcomeMessage, updateReq.Message.MessageID)
				}
			}
		}
	}
}
