package internal

import "sync"

type Pool struct {
	wg       sync.WaitGroup
	taskChan chan func()
}

func NewPool(workerNum int) *Pool {
	pool := &Pool{
		taskChan: make(chan func()),
	}

	for i := 0; i < workerNum; i++ {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()
			for task := range pool.taskChan {
				task()
			}
		}()
	}

	return pool
}

// 提交任务
func (p *Pool) Submit(task func()) {
	p.taskChan <- task
}

func (p *Pool) CloseAndWait() {
	close(p.taskChan)
	p.wg.Wait()
}
