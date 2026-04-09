package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	urlStr := "http://localhost:8080/api/fast"
	result := fetch(urlStr, 0)

	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
