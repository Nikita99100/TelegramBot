package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	output, err := outputTask(r)
	if err != nil {
		logs.Error(err, "Cant output tasks")
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		logs.Error(errors.Wrap(err, "Cant encode tasks and send to bot."))
	}
}
func createTask(w http.ResponseWriter, r *http.Request) {
	err := addTask(r)
	if err != nil {
		logs.Error(errors.Wrap(err, "Cant create task."))
	}
	response := Response{
		Status: "OK",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to encode response"))
	}
}
func doTask(w http.ResponseWriter, r *http.Request) {
	status, err := deleteTask(r)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to delete task"))
	}
	response := Response{
		Status: status,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to encode response"))
	}
}
