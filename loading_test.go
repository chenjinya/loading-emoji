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
		l.Loading(fmt.Sprintf(" trigger when next ticket, count down: %.2d", 10 - l.CountDownIndex() / (int(time.Second) / int(l.SleepDuration()))))
	})
	println("loading 1s")
	time.Sleep(10 * time.Second)
	l.Stop()
	println("loading done")
	time.Sleep(2 * time.Second)
}
