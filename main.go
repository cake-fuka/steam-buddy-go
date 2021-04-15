package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"ghe.corp.yahoo.co.jp/amasuda/mimir/di"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 10)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		defer close(sig)
		<-sig
		cancel()
	}()

	dc := di.NewContainer()

	// Initialize the components
	exporterServer := dc.ExporterServer()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(10 * time.Second) {
			// steamを見に行く処理のservice
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// healthのmetrics endpoint
	}()

	wg.Wait()
}
