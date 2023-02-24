package dwz

import (
	"fmt"
	"testing"
)

func TestGetUrl(t *testing.T) {
	Init()
	r, err := GetUrl("13")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func TestAddUrl(t *testing.T) {
	err := AddUrl("title", "https://example.com", "description")
	t.Error(err)
}
