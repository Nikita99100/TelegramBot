package main

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

var logs = logrus.New()

func main() {
	router := NewRouter()
	logs.Fatal(errors.Wrap(http.ListenAndServe(":8000", router), "Failed to start server"))
}
