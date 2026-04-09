### 用法说明

```
# 基本用法
gofetch --input urls.txt
# 从 stdin 读取
cat urls.txt | gofetch
# 选项
--input FILE URL 列表文件（不指定则读 stdin）
--concurrency N 并发请求数（默认值自定）
--timeout DURATION 单请求超时（默认值自定）
--output FILE 结果写入文件（不指定则写 stdout）
--format STRING 输出格式：json | jsonl | csv（默认自定）
--retry N 失败后重试次数（默认自定）
--verbose 打印进度信息到 stderr
```

### 第 5 节问题的决策

1. 输出顺序：并发抓取时，结果的输出顺序是否需要与输入 URL 的顺序保持一致？

 我的设计没有保证顺序性， 如果需要保证，可以预先创建一个结果切片 `results := make([]FetchResult, len(urls))`。并发 Worker 根据索引将结果填入切片的对应位置。待所有任务完成后，统一遍历切片输出。但是如果任务量极大，会导致内存占用飙升



2. 部分失败：某些 URL 请求失败时，程序如何继续运行？最终退出码如何设计？

每个抓取任务在独立的 Goroutine 中运行 。即使某个 Goroutine 因为 panic 或严重错误退出，也不应影响其他正在进行的 Worker。

程序退出：使用 `sync/atomic` 维护一个失败总数，根据失败数来执行os.Exit



3. 重试计时：--retry 重试失败请求时，latency_ms 应记录单次耗时还是累计耗时？

   我采用的是单次耗时，如果采用累计耗时，当一个请求在第 n次尝试成功时，其 latency_ms 可能会显示为一个巨大的数值(包含前面失败尝试的时间)，这在分析接口响应性能时会造成数据严重偏离。



4. URL 的校验：是否需要校验 URL 的合法性？

   可以预先校验， 避免无效的url发起请求造成资源浪费
