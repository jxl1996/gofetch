package main

import (
	"flag"
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

	flag.StringVar(&input, "input", "", "URL列表文件 (不指定则读 stdin)")
	flag.IntVar(&concurrency, "concurrency", 10, "并发请求数")
	flag.IntVar(&timeout, "timeout", 5, "单请求超时时间")
	flag.StringVar(&format, "format", "jsonl", "输出格式: json | jsonl | csv")
	flag.StringVar(&output, "output", "", "结果写入文件 (不指定则写 stdout)")
	flag.IntVar(&retry, "retry", 2, "失败后重试次数")
	flag.BoolVar(&verbose, "verbose", false, "打印进度信息到 stderr")

}
