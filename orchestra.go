package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/BlackChaosNL/Orchestra/cmd/api"
	"github.com/BlackChaosNL/Orchestra/cmd/web"
	"github.com/joho/godotenv"
)

/*
Orchestra is an visual orchestrator that allows you to setup services quickly.

This project is born out a need to quickly setup game services and to make it simple for friends and family, and to keep it in one binary without the setup hassle.
I would like it to be compliant to Pelican/Pterodactyl's Egg system so you can ingest those files and set up the game servers like a breeze.

TODO: Write usage.
INFO: Code organization: https://go.dev/doc/modules/layout
*/
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go api.StartAPIServer(&wg)
	go web.StartWebServer(&wg)

	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	_ = <-c // Wait for signal from OS.Signal

	api.StopAPIService()
	web.StopWebService()

	fmt.Print("Gracefully shutting down...")
}
