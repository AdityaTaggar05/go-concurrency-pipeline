package main

import "time"

type Response struct {
	URL string
	Status  int
	Latency time.Duration
	Error error
	NumTries int
}
