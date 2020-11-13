package main

import (
	"fmt"
	"github.com/pkg/errors"
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
func addTask(task ResponseTask) error {
	ourUser := FindUser(task.UserID)
	result := Task{
		Title: task.UserTask,
	}
	ourUser.Tasks = append(ourUser.Tasks, result)
	fmt.Println(ourUser)
	return nil
}
func outputTask(inputID ResponseID) ([]Task, error) {
	output, err := findTasks(inputID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to find tasks for current ID.")
	}
	return output, nil
}
func deleteTask(task ResponseTaskIndex) error {
	var err error
	ourUser := FindUser(task.UserID)
	ourUser.Tasks, err = removeElement(ourUser.Tasks, task.TaskIndex-1)
	if err != nil {
		return err
	}
	return nil
}
func updateTask(task ResponseTaskValue) error{
	var err error
	ourUser := FindUser(task.UserID)
	ourUser.Tasks, err = updateElement(ourUser.Tasks, task.TaskIndex-1, task.TaskValue)
	if err != nil {
		return err
	}
	return nil
}
func removeElement(s []Task, index int) ([]Task, error) {
	if index < len(s) {
		return append(s[:index], s[index+1:]...), nil
	}
	return s, errors.New("Index out of range")
}
func updateElement(s []Task, index int, value string) ([]Task, error) {
	if index < len(s) {
		s[index].Title = value
		return s, nil
	}
	return s, errors.New("Index out of range")
}
