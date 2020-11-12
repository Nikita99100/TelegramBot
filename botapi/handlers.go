package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pkg/errors"
	"strconv"
)

func AddTask(userID string, task string) (string, error) {
	request := ReqStruct{
		UserId: userID,
		Task:   task,
	}
	url := config.ServerUrl + ":" + config.Port + config.AddTaskUrl
	fmt.Println(request)
	err := MakeRequest("POST", url, request, nil)
	if err != nil {
		return "", errors.Wrap(err, "Add task failed")
	}
	return fmt.Sprintf("Added \"%s\" to your task list.\n", task), nil
}

func ListTasks(userID string) (string, error) {
	request := ReqStruct{
		UserId: userID,
	}
	url := config.ServerUrl + ":" + config.Port + config.ListTaskUrl
	var tasks []Task
	err := MakeRequest("GET", url, request, &tasks)
	if err != nil {
		return "", errors.Wrap(err, "Request list task failed")
	}
	return tasksToString(tasks), nil
}
func GetFile(chatId int64) (tgbotapi.DocumentConfig, error) {
	return tgbotapi.NewDocumentUpload(chatId, "botapi/static/file.txt"), nil
}
func DoTask(chatId string, taskIndex string) (string, error) {
	index, err := strconv.Atoi(taskIndex)
	if err != nil {
		return "", errors.Wrap(err, "Task index convert error")
	}
	request := ReqTaskIndex{
		UserID:    chatId,
		TaskIndex: index,
	}
	url := config.ServerUrl + ":" + config.Port + config.DoTaskUrl
	var response Response
	err = MakeRequest("DELETE", url, request, &response)
	if err != nil {
		return "", errors.Wrap(err, "Request delete task failed")
	}
	return response.Status, nil
}
