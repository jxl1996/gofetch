package internal

import (
	"io"
	"net/http"
	"time"
)

type FetchResult struct {
	URL        string    `json:"url"`
	StatusCode int       `json:"status_code"`
	LatencyMs  int64     `json:"latency_ms"`
	BodySize   int       `json:"body_size"`
	Error      error     `json:"error"`
	Timestamp  time.Time `json:"timestamp"`
	Attempt    int       `json:"attempt"`
}

type Fetcher struct {
	client *http.Client
	retry  int // 失败重试次数
}

func NewFetcher(retry int, timeout time.Duration) *Fetcher {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Fetcher{
		client: client,
		retry:  retry,
	}
}

func (f *Fetcher) FetchWithRetry(urlStr string) (result FetchResult) {
	for i := 0; i < f.retry; i++ {
		result = f.doHTTP(urlStr, i)
		if result.Error == nil {
			return result
		}

		if i < f.retry-1 {
			// 指数退避策略
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	// 返回最后一次的失败结果
	return
}

func (f *Fetcher) doHTTP(urlStr string, attempt int) (result FetchResult) {
	// 发起请求
	startTime := time.Now()
	resp, err := f.client.Get(urlStr)

	// 整理返回结构
	result.URL = urlStr
	result.LatencyMs = time.Since(startTime).Milliseconds()
	result.Timestamp = time.Now()
	result.Attempt = attempt + 1
	result.Error = err

	if err == nil {
		defer resp.Body.Close()
		result.StatusCode = resp.StatusCode

		if body, err := io.ReadAll(resp.Body); err == nil {
			result.BodySize = len(body)
		} else {
			result.Error = err
		}
	}

	return result
}
