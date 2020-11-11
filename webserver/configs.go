package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerUrl   string
	Port        string
	ListTaskUrl string
	AddTaskUrl  string
}

var config Config

func init() {
	configFile := viper.New()
	configFile.SetConfigFile("webserver/conf.toml")
	err := configFile.ReadInConfig()
	if err != nil {
		logs.Warn(err, "Cant read config")
	}
	errMar := configFile.Unmarshal(&config)
	if errMar != nil {
		logs.Warn(err, "Cant unmarshal config")
	}
}
