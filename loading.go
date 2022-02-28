package loji

import (
	"fmt"
	"time"
)

type LoadingEmoji struct {
	clocks []rune
	index  int
	status string
	signal chan string
	text   string
}

func NewLoading(clocks string) *LoadingEmoji {
	if clocks == "" {
		clocks = "🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚🕛"
	}
	return &LoadingEmoji{
		clocks: []rune(clocks),
		signal: make(chan string),
	}
}

func (l *LoadingEmoji) Loading(msg string) {
	go func() {
		l.loading(msg)
	}()
}

func (l *LoadingEmoji) Stop() {
	l.signal <- "stop"
}

func (l *LoadingEmoji) loading(msg string) {
	index := 0
	isStop := false
	go func() {
		for true {
			fmt.Printf("\r %s%s", string(l.clocks[index]), msg)
			index++
			if index >= len(l.clocks) {
				index = 0
			}
			time.Sleep(100 * time.Millisecond)
			if isStop {
				fmt.Println("")
				break
			}
		}
	}()

	select {
	case _ = <-l.signal:
		isStop = true
	}

}
