package main

import (
	"fmt"

	"github.com/fuel-/gideon/bot"
	"github.com/fuel-/gideon/config"
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
