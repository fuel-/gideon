package main

import (
	"fmt"

	"github.com/gideon/GoBot/bot"
	"github.com/gideon/GoBot/config"
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
