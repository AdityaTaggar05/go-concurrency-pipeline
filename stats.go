package main

import "time"

type Stats struct {
	MaxTime time.Duration
	MinTime time.Duration
	AvgTime time.Duration
	Count   int
}

func (s *Stats) Update(res Response) {
	s.Count++

	if res.Latency > s.MaxTime {
		s.MaxTime = res.Latency
	}

	if res.Latency < s.MinTime {
		s.MinTime = res.Latency
	}

	s.AvgTime = (time.Duration(s.Count - 1) * s.AvgTime + res.Latency) / time.Duration(s.Count)
}
