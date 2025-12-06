package main

import (
	"category-service/cmd/app"
	"log"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatalf("category-service not started: %s", err.Error())
	}
}
