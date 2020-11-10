package main

import (
	"github.com/spf13/viper"
)

var ServerUrl = "http://localhost:8000"
var ListTaskUrl = "/api/tasks/"
var AddTaskUrl = "/api/tasks/create/"

var configFile Config

func init(){
	config := viper.New()
	config.SetConfigFile("botapi/conf.toml")
	config.ReadInConfig()
	config.Unmarshal(&configFile)
}
