package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"tg.bot/bot"
	"tg.bot/cmd"
	"tg.bot/endpoint"
	"tg.bot/transport"
)

const (
	listenAddress = ":443"
	serverCRT     = "./server.crt"
	serverKEY     = "./server.key"
)

var (
	botToken       = bot.Token(os.Getenv("BOT_TOKEN"))
	webhookAddress = os.Getenv("WEBHOOK_ADDRESS")
)

func main() {
	var (
		logger log.Logger = cmd.MakeLogger()

		botService  *bot.Service  = bot.NewService(logger, botToken)
		endpointSet *endpoint.Set = endpoint.NewSet(botService)

		httpHandler http.Handler = transport.MakeHTTPHandler(
			*botService,
			*endpointSet,
			logger,
		)
	)

	level.Info(logger).Log("msg", "start")
	defer level.Info(logger).Log("msg", "stop")

	if err := http.ListenAndServeTLS(listenAddress, serverCRT, serverKEY, httpHandler); err != nil {
		level.Error(logger).Log(
			"msg", "failed to start server",
			"err", err,
		)
		os.Exit(1)
	}
}
