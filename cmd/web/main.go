package web

import (
	"sync"

	"github.com/kataras/golog"
)

func StartWebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	golog.Print("Web started")
}
