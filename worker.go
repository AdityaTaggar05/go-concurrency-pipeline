package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	maxRetries  = 3
	baseBackoff = 200 * time.Millisecond
	maxBackoff  = 5 * time.Second
)

func GenerateWorkers(ctx context.Context, jobs <-chan string, results chan<- Response) {
	numWorkers := 8

	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Go(func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()

			dispatch(ctx, i, jobs, results)
		})
	}

	wg.Wait()
	close(results)
}

func dispatch(ctx context.Context, id int, jobs <-chan string, results chan<- Response) {
	fmt.Println("Worker", id, "is up and running! Starting in 2 seconds...")
	time.Sleep(time.Second * 2)

	// HTTP Client
	client := &http.Client{Timeout: time.Second * 10}

	// Rate Limiting
	limiter := rate.NewLimiter(10, 4)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping Worker %d...\n", id)
			return
		case url, ok := <-jobs:
			if !ok {
				return
			}

			if err := limiter.Wait(ctx); err != nil {
				fmt.Printf("Stopping Worker %d...\n", id)
				return
			}

			for i := range maxRetries {
				if ctx.Err() != nil {
					fmt.Printf("Stopping Worker %d...\n", id)
					return
				}

				start := time.Now()
				status, err := sendRequest(client, url)

				res := Response{
					URL:     url,
					Latency: time.Since(start),
				}

				if err != nil {
					res.Error = err
				} else {
					res.Status = status
				}

				results <- res

				if isRetryable(status, err) {
					dur := backoff(i)
					fmt.Printf("Retry %d for url (%s) running in %v\n", i+2, url, dur)

					select {
					case <-ctx.Done():
						fmt.Printf("Stopping Worker %d...\n", id)
						return
					case <-time.After(dur):
					}
				} else {
					break
				}
			}
		}
	}
}

func sendRequest(client *http.Client, url string) (int, error) {
	res, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

func isRetryable(status int, err error) bool {
	if err != nil {
		return true
	}

	if status >= 500 {
		return true
	}

	return false
}

func backoff(attempt int) time.Duration {
	d := min(baseBackoff*(1<<attempt), maxBackoff)
	jitter := time.Duration(rand.Int63n(int64(d / 3)))

	return 2*d/3 + jitter
}
