package main

import (
	"fmt"
	"time"
)

type Stats struct {
	MaxTime time.Duration
	MinTime time.Duration
	AvgTime time.Duration
	Count   int
}

func NewStats() *Stats {
	return &Stats{
		MinTime: time.Hour,
	}
}

func (s *Stats) Update(res Response) {
	s.Count += res.NumTries

	if res.Latency > s.MaxTime {
		s.MaxTime = res.Latency
	}

	if res.Latency < s.MinTime {
		s.MinTime = res.Latency
	}

	s.AvgTime = (time.Duration(s.Count-1)*s.AvgTime + res.Latency) / time.Duration(s.Count)
}

func (s *Stats) String() string {
	return fmt.Sprintf(
		`Stats for %d requests made
-----------------------------------------------------
Max Time: %s
Min Time: %s
Avg Time: %s
-----------------------------------------------------`,
		s.Count, s.MaxTime, s.MinTime, s.AvgTime)
}
