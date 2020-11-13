package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var logs = logrus.New()

func main() {
	bot, err := tgbotapi.NewBotAPI("1158930916:AAH1CZKaNxGhDJrNg6w0SakIH-48cH4Jl9o")
	if err != nil {
		logs.Fatal(errors.Wrap(err, "Failed connect to bot"))
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "Failed to get update"))
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg, err := MsgHandler(update.Message)
		if err != nil {
			logs.Error(errors.Wrap(err, "Failed to handle message"))
		}
		bot.Send(msg)
	}
}
