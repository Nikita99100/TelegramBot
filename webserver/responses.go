package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type ResponseTask struct {
	UserID   string `json:"user_id"`
	UserTask string `json:"user_task"`
}
type ResponseTaskIndex struct {
	UserID    string `json:"user_id"`
	TaskIndex int    `json:"task_index"`
}
type ResponseID struct {
	UserID string `json:"user_id"`
}

func unmarshalRequest(r *http.Request, payLoad interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "Failed to read request body.")
	}
	err = json.Unmarshal(body, &payLoad)
	if err != nil {
		return errors.Wrap(err, "Unmarshal error")
	}
	return nil
}
