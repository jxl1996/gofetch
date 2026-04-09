package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jxl1996/gofetch/internal"
	"github.com/jxl1996/gofetch/utils"
	"io"
	"os"
	"time"
)

func main() {
	var (
		input       string
		concurrency int
		timeout     int
		output      string
		format      string
		retry       int
		verbose     bool
	)

	flag.StringVar(&input, "input", "urls.txt", "URL列表文件 (不指定则读 stdin)")
	flag.IntVar(&concurrency, "concurrency", 10, "并发请求数")
	flag.IntVar(&timeout, "timeout", 5, "单请求超时时间")
	flag.StringVar(&format, "format", "jsonl", "输出格式: json | jsonl | csv")
	flag.StringVar(&output, "output", "", "结果写入文件 (不指定则写 stdout)")
	flag.IntVar(&retry, "retry", 2, "失败后重试次数")
	flag.BoolVar(&verbose, "verbose", false, "打印进度信息到 stderr")

	//  准备输入源
	var reader io.Reader
	if input != "" {
		f, err := os.Open(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "无法打开输入文件: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		reader = f
	} else {
		reader = os.Stdin
	}

	// 收集结果
	resultChan := make(chan internal.FetchResult, concurrency)
	done := make(chan struct{})
	go func() {
		writeResult(resultChan)
		done <- struct{}{}
	}()

	// 创建任务池
	pool := internal.NewPool(concurrency)
	// 读取每一行 并提交到任务池
	fetcher := internal.NewFetcher(retry, time.Duration(timeout)*time.Second)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if urlStr := scanner.Text(); utils.IsValidURL(urlStr) {
			pool.Submit(func() {
				res := fetcher.FetchWithRetry(urlStr)
				resultChan <- res
			})
		}
	}
	// 等待fetch任务全部执行完毕
	pool.CloseAndWait()
	close(resultChan)

	// 等待收集结果完毕
	<-done
}

// 输出结果
func writeResult(resultChan chan internal.FetchResult) {
	for result := range resultChan {
		fmt.Println(result)
	}
}
