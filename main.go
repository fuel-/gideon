package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fuel-/gideon/bot"
	"github.com/fuel-/gideon/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)

	<-make(chan struct{})
	return
}
