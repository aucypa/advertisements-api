package main

import (
	"log"
	"users-service/cmd/app"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatalf("users-service not started: %s", err.Error())
	}
}
