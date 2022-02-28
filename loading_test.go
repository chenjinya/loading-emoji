package loji

import (
	"testing"
	"time"
)

func TestNewLoading(t *testing.T) {
	l := New()
	l.Loading("loading...")
	l.Loading("loading...2")
	l.Loading("loading...33")
	println("loading 1s")
	time.Sleep(1 * time.Second)
	l.Stop()
	time.Sleep(2 * time.Second)
}
