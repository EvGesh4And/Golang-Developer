package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0] // <-or(ch) == <-ch
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[:3], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			defer close(c)
			<-time.After(after)
		}()
		return c
	}

	var slice []<-chan struct{}

	for i := 0; i < 10; i++ {
		slice = append(slice, sig(10*time.Second))
	}
	slice = append(slice, sig(time.Millisecond*1))

	start := time.Now()

	<-or3(slice...)
	// <-or2(true, slice...)

	fmt.Printf("done after %v", time.Since(start))
}
