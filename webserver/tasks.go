package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func findTasks(inputID ResponseID) (output []Task, err error) {
	for _, user := range users {
		if user.ID == inputID.UserID {
			for _, task := range user.Tasks {
				output = append(output, task)
			}
		}
	}
	return output, nil
}
func addTask(r *http.Request) error {
	var task ResponseTask
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "Failed to read request body.")
	}
	err = json.Unmarshal(body, &task)
	if err != nil {
		return errors.Wrap(err, "Cant decode current task.")
	}
	ourUser := FindUser(task.UserID)
	result := Task{
		ID:    strconv.Itoa(len(ourUser.Tasks) + 1),
		Title: task.UserTask,
	}
	ourUser.Tasks = append(ourUser.Tasks, result)
	users[FindUserIndex(ourUser.ID)] = ourUser
	return nil
}
func outputTask(r *http.Request) ([]Task, error) {
	var inputID ResponseID
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read request body.")
	}
	err = json.Unmarshal(body, &inputID)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal error")
	}
	output, err := findTasks(inputID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to find tasks for current ID.")
	}
	return output, nil
}
