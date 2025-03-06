package main

import (
	"time"
)

// go run -race main.go
func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}
	time.Sleep(time.Second)
}
