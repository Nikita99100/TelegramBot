package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pkg/errors"
)

func AddTask(userID string, task string) (string, error) {
	request, err := NewReq(userID, task)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create new req")
	}
	url := config.ServerUrl + ":" + config.Port + config.AddTaskUrl
	err = MakeRequest("POST", url, request, nil)
	if err != nil {
		return "", errors.Wrap(err, "Add task failed")
	}
	return fmt.Sprintf("Added \"%s\" to your task list.\n", task), nil
}

func ListTasks(userID string) (string, error) {
	request, err := NewReq(userID, "")
	if err != nil {
		return "", errors.Wrap(err, "Failed to create new req")
	}
	url := config.ServerUrl + ":" + config.Port + config.ListTaskUrl
	var tasks []Task
	err = MakeRequest("GET", url, request, &tasks)
	if err != nil {
		return "", errors.Wrap(err, "Request list task failed")
	}
	return tasksToString(tasks), nil
}

func DoTask(chatId string, taskIndex string) error {
	request, err := NewReqTaskIndex(chatId, taskIndex)
	if err != nil {
		return errors.Wrap(err, "Failed to create new req")
	}
	url := config.ServerUrl + ":" + config.Port + config.DoTaskUrl
	err = MakeRequest("DELETE", url, request, nil)
	if err != nil {
		return errors.Wrap(err, "Request delete task failed")
	}
	return nil
}
func EditTask(chatId string, taskData string) error {
	index, value, err := spaceParse(taskData)
	if err != nil {
		logs.Warn(err)
	}
	request := ReqTaskValue{
		UserID:    chatId,
		TaskIndex: index,
		TaskValue: value,
	}
	url := config.ServerUrl + ":" + config.Port + config.EditTaskUrl
	err = MakeRequest("PUT", url, request, nil)
	if err != nil {
		return errors.Wrap(err, "Request edit task failed")
	}
	return nil
}

func GetFile(chatId int64) (tgbotapi.DocumentConfig, error) {
	return tgbotapi.NewDocumentUpload(chatId, "botapi/static/file.txt"), nil
}
