package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	var inputID ResponseID
	err := unmarshalRequest(r, &inputID)
	if err != nil {
		logs.Error("Failed to unmarshal request")
		w.WriteHeader(http.StatusBadRequest)
	}
	output, err := outputTask(inputID)
	if err != nil {
		logs.Error(err, "Cant output tasks")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(output) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		logs.Error(errors.Wrap(err, "Cant encode tasks and send to bot."))
		w.WriteHeader(http.StatusBadRequest)
	}
}
func createTask(w http.ResponseWriter, r *http.Request) {
	var task ResponseTask
	err := unmarshalRequest(r, &task)
	if err != nil {
		logs.Error("Failed to unmarshal request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = addTask(task)
	var response Response
	if err != nil {
		logs.Error(errors.Wrap(err, "Cant create task."))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Something bad happens :("))
		response = Response{
			Status: "NOT OK",
		}
		return
	}
	response = Response{
		Status: "OK",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to encode response"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Something bad happens :("))
	}
}
func doTask(w http.ResponseWriter, r *http.Request) {
	var task ResponseTaskIndex
	err := unmarshalRequest(r, &task)
	if err != nil {
		logs.Error(err, "Failed to unmarshal request")
	}
	err = deleteTask(task)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to delete task"))
		w.WriteHeader(http.StatusBadRequest)
	}

}
func editTask(w http.ResponseWriter, r *http.Request) {
	var task ResponseTaskValue
	err := unmarshalRequest(r, &task)
	if err != nil {
		logs.Error(err, "Failed to unmarshal request")
		w.WriteHeader(http.StatusBadRequest)
	}
	err = updateTask(task)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to delete task"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
