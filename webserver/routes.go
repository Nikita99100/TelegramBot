package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"ListTasks",
		"GET",
		"/api/tasks/",
		getTasks,
	},
	Route{
		"CreateTask",
		"POST",
		"/api/tasks/create/",
		createTask,
	},
	Route{
		"DeleteTask",
		"DELETE",
		"/api/tasks/do/",
		doTask,
	},
	Route{
		Name:        "EditTask",
		Method:      "PUT",
		Pattern:     "/api/tasks/edit",
		HandlerFunc: editTask,
	},
}
