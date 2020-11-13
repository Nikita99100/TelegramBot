package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func MsgHandler(msgReceive *tgbotapi.Message) (tgbotapi.Chattable, error) {
	msgTxt := CmdClipper(msgReceive)
	msgSend, err := CmdHandler(msgReceive.Command(), msgTxt, msgReceive.Chat)
	if err != nil {
		return tgbotapi.MessageConfig{}, errors.Wrap(err, "Failed to handle command")
	}
	return msgSend, nil
}

func CmdHandler(command string, text string, chat *tgbotapi.Chat) (tgbotapi.Chattable, error) {
	var msg tgbotapi.Chattable
	var err error
	switch command {
	case "do":
		msg, err = DoTaskMessage(chat.ID, text)
	case "get":
		msg, err = GetFileMessage(chat.ID)
		if err != nil {
			return tgbotapi.MessageConfig{}, errors.Wrap(err, "Create GetFileMessage error")
		}
	case "add":
		msg, err = AddTaskMessage(chat.ID, text)
		if err != nil {
			return tgbotapi.MessageConfig{}, errors.Wrap(err, "Create AddTaskMessage error")
		}
	case "list":
		msg, err = ListTaskMessage(chat.ID)
		if err != nil {
			return tgbotapi.MessageConfig{}, errors.Wrap(err, "Create ListTaskMessage error")
		}
	case "start":
		msg, err = StartMessage(chat)
		if err != nil {
			return tgbotapi.MessageConfig{}, errors.Wrap(err, "Start command error")
		}
	default:
		msg = tgbotapi.NewMessage(chat.ID, "Такой комманды нет")
	}
	return msg, nil
}

func CmdClipper(msg *tgbotapi.Message) string {
	text := msg.Text
	if msg.Command() != "" {
		text = strings.Join(strings.Split(text, " ")[1:], " ")
	}
	return text
}
func spaceParse(text string) (int, string, error) {
	s := strings.Split(text, " ")
	index, err := strconv.Atoi(s[0])
	if err != nil {
		logs.Error(err, "Cant convert index")
		return 0, "", err
	}
	if len(s[1]) == 0 {
		logs.Error("Task value was empty")
		return 0, "", errors.New("Task value was empty")
	}
	return index, s[1], nil
}
