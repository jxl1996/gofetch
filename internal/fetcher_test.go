package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchWithRetry(t *testing.T) {
	f := NewFetcher(3, time.Second*3)
	res := f.FetchWithRetry("http://localhost:8080/api/slow")
	fmt.Printf("%#v\n", res)
}
