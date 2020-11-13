package main

import "fmt"

type Task struct {
	Task string `json:"title"`
}

func tasksToString(structs []Task) string {
	s := ""
	for i, v := range structs {
		s += fmt.Sprintf("%d. %s\n", i+1, v.Task)
	}
	return s
}
