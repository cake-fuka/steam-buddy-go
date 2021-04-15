package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/cake-fuka/steam-buddy-go/di"
)

func main() {
	dc := di.NewContainer()

	// Initialize the components
	s := dc.Service()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(10 * time.Second) {
			// steamを見に行く処理のservice
			err := s.ObservSteam()
			if err != nil {
				fmt.Errorf("困った")
			}
			fmt.Println("steam check")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// healthのmetrics endpoint
		fmt.Println("health check standby")
	}()

	wg.Wait()
}
