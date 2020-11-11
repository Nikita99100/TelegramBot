package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Task struct {
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
	err := unmarshalRequest(r, &task)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarhal request")
	}
	ourUser := FindUser(task.UserID)
	result := Task{
		Title: task.UserTask,
	}
	ourUser.Tasks = append(ourUser.Tasks, result)
	fmt.Println(ourUser)
	return nil
}
func outputTask(r *http.Request) ([]Task, error) {
	var inputID ResponseID
	err := unmarshalRequest(r, &inputID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarhal request")
	}
	output, err := findTasks(inputID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to find tasks for current ID.")
	}
	return output, nil
}
func deleteTask(r *http.Request) (string, error) {
	var task ResponseTaskIndex
	fmt.Println("delete")
	err := unmarshalRequest(r, &task)
	if err != nil {
		return "FAILED", errors.Wrap(err, "Failed to unmarhal request")
	}
	ourUser := FindUser(task.UserID)
	ourUser.Tasks, err = removeElement(ourUser.Tasks, task.TaskIndex-1)
	if err != nil {
		return "FAILED", nil
	}
	return "OK", nil
}
func removeElement(s []Task, index int) ([]Task, error) {
	if index < len(s) {
		return append(s[:index], s[index+1:]...), nil
	}
	return s, errors.New("Index out of range")
}
