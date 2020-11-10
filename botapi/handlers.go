package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func AddTask(userID string, task string) (string, error) {
	request := ReqStruct{
		UserId: userID,
		Task:   task,
	}
	url := ServerUrl + AddTaskUrl
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
	url := ServerUrl + ListTaskUrl
	var tasks []Task
	err := MakeRequest("GET", url, request, &tasks)
	if err != nil {
		return "", errors.Wrap(err, "Request list task failed")
	}
	return structsToString(tasks), nil
}
func GetFile() {

}
