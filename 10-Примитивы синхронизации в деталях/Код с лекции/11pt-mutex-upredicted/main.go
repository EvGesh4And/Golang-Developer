package main

import (
	"fmt"
	"sync"
)

// go run -race main.go
func main() {
	wg := sync.WaitGroup{}
	text := "old text"
	wg.Add(2)
	go func() {
		text = "hello world"
		wg.Done()
	}()

	go func() {
		fmt.Println(text)
		wg.Done()
	}()
	wg.Wait()
}
