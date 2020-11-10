package main

import (
	"github.com/spf13/viper"
)

var config Config

func init(){
	configFile := viper.New()
	configFile.SetConfigFile("botapi/conf.toml")
	configFile.ReadInConfig()
	configFile.Unmarshal(&config)
}
