package main

import "fmt"

type Task struct {
	Task string `json:"title"`
}
type ReqStruct struct {
	UserId string `json:"user_id"`
	Task   string `json:"user_task"`
}
type ReqTaskIndex struct {
	UserID    string `json:"user_id"`
	TaskIndex int    `json:"task_index"`
}
type Response struct {
	Status string `json:"status"`
}

func tasksToString(structs []Task) string {
	s := ""
	for i, v := range structs {
		s += fmt.Sprintf("%d. %s\n", i+1, v.Task)
	}
	return s
}
