package main

import (
	"fmt"
	"gideon/GoBot/bot"
	"gideon/GoBot/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
