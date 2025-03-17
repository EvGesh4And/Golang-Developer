package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intStream := make(chan int)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case i := <-intStream:
				fmt.Println(i)
			case <-done:
				return
			}
		}
	}()

	go func() {
		time.Sleep(time.Second*6 + time.Second*time.Duration(rand.Intn(3)))
		close(done)
	}()

	for _, i := range []int{1, 2, 3, 4, 5} {
		select {
		case <-done:
			return
		case intStream <- i:
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		}
	}
}
