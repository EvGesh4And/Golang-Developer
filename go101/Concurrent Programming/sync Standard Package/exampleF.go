package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Hi")
		m.Unlock() // make a notification
	}()
	m.Lock() // wait to be notified
	fmt.Println("Bye")
}
