package main

import "fmt"

type Task struct {
	Id   string `json:"id"`
	Task string `json:"title"`
}
type ReqStruct struct {
	UserId string `json:"user_id"`
	Task   string `json:"user_task"`
}
type Config struct {
	ServerUrl           string
	ListTaskUrl   		string
	AddTaskUrl 		  	string
}

func structsToString(structs []Task) string {
	s := ""
	for _, v := range structs {
		s += fmt.Sprintf("%s. %s\n", v.Id, v.Task)
	}
	return s
}
