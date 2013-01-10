package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Page_Size       int
	DB_Name         string
	Job_Collection  string
	Info_Collection string
	MongoDB_Server  string
	Server_Port     string
}

var Configuration Config

func InitConfig() Config {
	var c Config
	content, _ := ioutil.ReadFile("config.json")
	_ = json.Unmarshal(content, &c)

	return c
}
