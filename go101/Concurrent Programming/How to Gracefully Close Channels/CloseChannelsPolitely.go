package main

import (
	"fmt"
	"sync"
)

type MyChannel struct {
	c    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{c: make(chan int)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.c)
	})
}

func SafeClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

func main() {
	mych := NewMyChannel()

	mych.SafeClose()

	mych.SafeClose()

	fmt.Println(SafeClose(mych.c))
}
