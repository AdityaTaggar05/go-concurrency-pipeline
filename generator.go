package main

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"os"
)

func GenerateJobs(ctx context.Context, jobs chan<- string, path string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}

		close(jobs)
	}()

	// File Handling
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("[ERR] failed to open file: %s", err)
		return err
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	// Line-by-Line Scanning
	for reader.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			urlPath := reader.Text()

			if _, err := url.ParseRequestURI(urlPath); err != nil {
				fmt.Printf("❌ invalid url format: %v\n", err)
			} else {
				jobs <- urlPath
			}

		}
	}

	if err := reader.Err(); err != nil {
		fmt.Printf("[ERR] scan error: %v\n", err)
	}

	return nil
}
