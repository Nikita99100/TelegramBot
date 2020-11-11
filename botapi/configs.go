package main

import "github.com/spf13/viper"

var ServerUrl = "http://localhost:8000"
var ListTaskUrl = "/api/tasks/"
var AddTaskUrl = "/api/tasks/create/"

var configFile Config

func configParse() {
	config := viper.New()
	config.SetConfigFile("conf.toml")
	config.ReadInConfig()
	config.Unmarshal(&configFile)
}
