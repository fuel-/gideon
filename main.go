package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fuel-/gideon/bot"
	"github.com/fuel-/gideon/config"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
