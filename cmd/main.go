package main

import (
	"log"

	"betfate.com/app"
)

func main() {

	server := app.New()
	if err := server.Start(); err != nil {
		log.Fatal("can't run the server")
	}

}
