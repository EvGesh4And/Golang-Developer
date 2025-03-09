package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// The capacity of the signal channel can
	// also be one. If this is true, then a
	// value must be sent to the channel before
	// creating the following goroutine.

	go func() {
		fmt.Print("Hello")
		// Simulate a workload.
		time.Sleep(time.Second * 2)

		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		<-done
	}()

	// Blocked here, wait for a notification.
	done <- struct{}{}
	fmt.Println(" world!")
}
