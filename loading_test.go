package loji

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLoading(t *testing.T) {
	l := New()
	l.Loading("loading...")
	l.Loading("loading...2")
	l.Loading("loading...33")
	l.NextTick(func() {
		l.Loading(fmt.Sprintf(" trigger when next ticket, index: %d", l.CountDownIndex()))
	})
	println("loading 1s")
	time.Sleep(1 * time.Second)
	l.Stop()
	println("loading done")
	time.Sleep(2 * time.Second)
}
