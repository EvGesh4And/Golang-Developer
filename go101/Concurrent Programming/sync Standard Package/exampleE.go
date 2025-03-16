package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.RWMutex
	go func() {
		m.RLock()
		fmt.Print("a")
		time.Sleep(time.Second)
		m.RUnlock()
	}()
	go func() {
		time.Sleep(time.Second * 1 / 4)
		m.Lock()
		fmt.Print("b")
		time.Sleep(time.Second)
		m.Unlock()
	}()
	go func() {
		time.Sleep(time.Second * 2 / 4)
		m.Lock()
		fmt.Print("c")
		m.Unlock()
	}()
	go func() {
		time.Sleep(time.Second * 3 / 4)
		m.RLock()
		fmt.Print("d")
		m.RUnlock()
	}()
	time.Sleep(time.Second * 3)
	fmt.Println()
}
