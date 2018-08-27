package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	//Token is our bots token id
	Token string
	//BotPrefix is an identifier that our bot will use to know when it is being spoken to
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

//ReadConfig reads our config file and returns values set in said file
func ReadConfig() error {
	fmt.Println("Reading from config file . . .")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
