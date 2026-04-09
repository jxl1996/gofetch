package internal

import "time"

type FetchResult struct {
	URL        string    `json:"url"`
	StatusCode int       `json:"status_code"`
	LatencyMs  int64     `json:"latency_ms"`
	BodySize   int       `json:"body_size"`
	Error      error     `json:"error"`
	Timestamp  time.Time `json:"timestamp"`
	Attempt    int       `json:"attempt"`
}
