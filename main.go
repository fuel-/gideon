package main

import (
	"fmt"
	"os"

	"github.com/fuel-/gideon/bot"
	"github.com/fuel-/gideon/config"
)

func main() {
	fmt.Println(os.Getwd())
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
