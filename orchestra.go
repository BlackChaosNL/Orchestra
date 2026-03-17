package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/BlackChaosNL/Orchestra/cmd/api"
	"github.com/BlackChaosNL/Orchestra/cmd/web"
	"github.com/kataras/golog"
)

/*
Orchestra is an visual orchestrator that allows you to setup services quickly.

This project is born out a need to quickly setup game services and to make it simple for friends and family, and to keep it in one binary without the setup hassle.
I would like it to be compliant to Pelican/Pterodactyl's Egg system so you can ingest those files and set up the game servers like a breeze.

TODO: Write usage.
INFO: Code organization: https://go.dev/doc/modules/layout
*/
func main() {
	var wg sync.WaitGroup
	golog.Install(log.New(os.Stdout, "", 0))

	wg.Add(2)

	go api.StartAPIServer(&wg)
	go web.StartWebServer(&wg)

	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c // Wait for signal from OS.Signal
	api.StopAPIService()
	fmt.Println("Gracefully shutting down...")
}
