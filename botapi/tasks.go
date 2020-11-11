package main

import "fmt"

type Task struct {
	Id   string `json:"id"`
	Task string `json:"title"`
}

func tasksToString(structs []Task) string {
	s := ""
	for _, v := range structs {
		s += fmt.Sprintf("%s. %s\n", v.Id, v.Task)
	}
	return s
}
