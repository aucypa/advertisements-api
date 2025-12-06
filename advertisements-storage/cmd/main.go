package main

import (
	"log"

	"advertisement-storage/cmd/app"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatalf("advertisements-storage not started: %s", err.Error())
	}
}
