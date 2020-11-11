package main

import (
	"fmt"
	"github.com/pkg/errors"
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
	err := unmarshalRequest(r, &task)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarhal request")
	}
	ourUser := FindUser(task.UserID)
	result := Task{
		ID:    strconv.Itoa(len(ourUser.Tasks) + 1),
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
func deleteTask(r *http.Request) error {
	var task ResponseTaskIndex
	err := unmarshalRequest(r, &task)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarhal request")
	}
	return nil
	ourUser := FindUser(task.UserID)
	fmt.Println(ourUser.Tasks)
}
func RemoveTask(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
