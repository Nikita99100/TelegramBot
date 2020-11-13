package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type ReqStruct struct {
	UserId string `json:"user_id"`
	Task   string `json:"user_task"`
}
type ReqTaskIndex struct {
	UserID    string `json:"user_id"`
	TaskIndex int    `json:"task_index"`
}
type ReqTaskValue struct {
	UserID    string `json:"user_id"`
	TaskIndex int    `json:"task_index"`
	TaskValue string `json:"task_value"`
}
type Response struct {
	Status string `json:"status"`
}

func MakeRequest(method string, url string, payload, response interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal a payload")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return errors.Wrap(err, "failed to create an http request")
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to make a %s request", method))
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Errorf("%s request to %s failed with status: %d", method, url, resp.StatusCode)
		}
		return errors.Errorf("%s request to %s failed with status: %d and body: %s", method, url, resp.StatusCode, string(body))
	}
	if response != nil {
		return errors.Wrap(json.NewDecoder(resp.Body).Decode(response), "Failed to decode response")
	}
	return nil
}
