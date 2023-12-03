package main

import (
	"log"
	"os"
	"os/signal"

	"karmapi/app"
)

func main() {

	ch := make(chan os.Signal, 1)
	server := app.New()
	go func() {
		if err := server.Start(); err != nil {
			log.Fatal("can't run the server", err)
		}
	}()

	signal.Notify(ch, os.Interrupt)
	<-ch

}
