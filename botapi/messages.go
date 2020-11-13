package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pkg/errors"
	"strconv"
)

func StartMessage(chat *tgbotapi.Chat) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(chat.ID, fmt.Sprintf(`Привет @%s! Я тут <em>слежу</em> за порядком. Веди себя хорошо.`,
		chat.FirstName))
	button := tgbotapi.NewKeyboardButton("/list")
	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{button})
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Html"
	return msg, nil
}
func GetFileMessage(chatId int64) (tgbotapi.Chattable, error) {
	file, err := GetFile(chatId)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func AddTaskMessage(chatId int64, task string) (tgbotapi.MessageConfig, error) {
	msgTxt, err := AddTask(strconv.FormatInt(chatId, 10), task)
	if err != nil {
		return tgbotapi.MessageConfig{}, errors.Wrap(err, "Failed AddTask command")
	}
	msg := tgbotapi.NewMessage(chatId, msgTxt)
	return msg, nil
}
func ListTaskMessage(chatId int64) (tgbotapi.MessageConfig, error) {
	tasks, err := ListTasks(strconv.FormatInt(chatId, 10))
	if err != nil {
		return tgbotapi.MessageConfig{}, errors.Wrap(err, "Failed ListTask command")
	}
	msg := tgbotapi.NewMessage(chatId, tasks)
	return msg, nil
}

func DoTaskMessage(chatId int64, taskIndex string) (tgbotapi.Chattable, error) {
	err := DoTask(strconv.FormatInt(chatId, 10), taskIndex)
	fmt.Println(err)
	if err != nil {
		return tgbotapi.NewVoiceUpload(chatId, "botapi/static/sound.mp3"), errors.Wrap(err, "Failed to do task")
	}
	return tgbotapi.NewMessage(chatId, "Task marked assss completed!!"), nil
}
func EditTaskMessage(chatId int64, taskIndex string) (tgbotapi.Chattable, error) {
	err := EditTask(strconv.FormatInt(chatId, 10), taskIndex)
	if err != nil {
		return tgbotapi.NewVoiceUpload(chatId, "botapi/static/sound.mp3"), errors.Wrap(err, "Failed to edit task")
	}
	return tgbotapi.NewMessage(chatId, "Task was editetssss!!"), nil
}
