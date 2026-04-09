package internal

import (
	"io"
	"net/http"
	"time"
)

func fetch(urlStr string, attempt int) (result FetchResult) {
	client := &http.Client{}

	// 发起请求
	startTime := time.Now()
	resp, err := client.Get(urlStr)

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
