package main

import (
	"fmt"

	"github.com/gideon/bot"
	"github.com/gideon/config"
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
