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
		logs.Error(errors.Wrap(err, "Cant decode current task."))
	}
}
func doTask(w http.ResponseWriter, r *http.Request) {

}
