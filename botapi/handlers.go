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

func DoTask(chatId string, taskIndex string) (string, error) {
	request, err := NewReqTaskIndex(chatId, taskIndex)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create new req")
	}
	url := config.ServerUrl + ":" + config.Port + config.DoTaskUrl
	var response Response
	err = MakeRequest("DELETE", url, request, &response)
	if err != nil {
		return "", errors.Wrap(err, "Request delete task failed")
	}
	return response.Status, nil
}

func GetFile(chatId int64) (tgbotapi.DocumentConfig, error) {
	return tgbotapi.NewDocumentUpload(chatId, "botapi/static/file.txt"), nil
}
