package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	path := flag.String("file", "urls.txt", "The file containing urls")
	flag.Parse()

	numJobs := 10

	jobs := make(chan string, numJobs)
	results := make(chan Response, numJobs)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	start := time.Now()

	go func() {
		if err := GenerateJobs(ctx, jobs, *path); err != nil {
			fmt.Println("Stopping job generation...")
		}
	}()
	go GenerateWorkers(ctx, jobs, results)

	for res := range results {
		if res.Error != nil {
			fmt.Printf("❌ %s: %v (%s)\n", res.URL, res.Error, res.Latency)
		} else {
			fmt.Printf("✅ %s: %d (%s)\n", res.URL, res.Status, res.Latency)
		}
	}

	fmt.Printf("\n-----------------------------------------------------\nTotal Time Elapsed: %s\n", time.Since(start))
}
