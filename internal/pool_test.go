package internal

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	pool := NewPool(10)
	for i := 0; i < 100; i++ {
		j := i

		pool.Submit(func() {
			fmt.Println("task", j)
		})
	}
	pool.CloseAndWait()
}
