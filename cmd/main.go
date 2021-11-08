package main

import (
	"flag"
	"log"

	"betfate.com/app"
)

var configPath = "configs/server.yaml"

func main() {

	flag.Parse()
	config := app.NewConfig(configPath)

	server := app.New(config)
	if err := server.Start(); err != nil {
		log.Fatal("can't run the server")
	}

}
