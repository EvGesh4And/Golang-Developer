package main

import (
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

func main() {
	mych := NewMyChannel()
	mych.SafeClose()
	mych.SafeClose()
}
