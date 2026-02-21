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
	// Initialize CMD Argument
	path := *flag.String("file", "urls.txt", "The file containing urls")
	flag.String("o", "results.txt", "The file containing the results")

	flag.Int("retries", 3, "Number of retries per request")
	flag.Int("workers", 8, "Numer of workers to run concurrently")
	numJobs := *flag.Int("jobs", 10, "Number of jobs to handle concurrently")
	flag.Parse()

	jobs := make(chan string, numJobs)
	results := make(chan Response, numJobs)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	start := time.Now()

	go func() {
		if err := GenerateJobs(ctx, jobs, path); err != nil {
			fmt.Println("Stopping job generation...")
		}
	}()
	go GenerateWorkers(ctx, jobs, results)

	// Initialize Stats
	stats := Stats{}

	for res := range results {
		stats.Update(res)

		if res.Error != nil {
			fmt.Printf("❌ %s: %v (%s)\n", res.URL, res.Error, res.Latency)
		} else {
			fmt.Printf("✅ %s: %d (%s)\n", res.URL, res.Status, res.Latency)
		}
	}

	fmt.Printf("\n-----------------------------------------------------\nTotal Time Elapsed: %s\nStats: %v\n", time.Since(start), stats)
}
