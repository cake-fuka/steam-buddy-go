package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/robfig/cron/v3"

	"github.com/cake-fuka/steam-buddy-go/di"
)

func main() {
	dc := di.NewContainer()

	// Initialize the components
	s := dc.Service()

	c := cron.New()
	c.AddFunc("@every 10s", func() {
		// steamを見に行く処理のservice
		err := s.RecentCheck()
		if err != nil {
			fmt.Errorf("困った")
		}
	})
	c.AddFunc("0 0 10 * * 0", func() {
		// 最近のゲームのプレイ時間をslackに送信する処理
		fmt.Println("output message every sunday at 10:00")
		err := s.WeeklyCheck()
		if err != nil {
			fmt.Errorf("困った")
		}
	})
	c.Start()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer c.Stop()

		// healthのmetrics endpoint
		fmt.Println("health check standby")
		http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {})
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf(err.Error())
		}
	}()

	wg.Wait()
}
