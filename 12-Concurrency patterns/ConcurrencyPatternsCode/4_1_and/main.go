package main

import (
	"fmt"
	"time"
)

func sig(t time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		<-time.After(t)
	}()
	return ch
}

func main() {

	var slice []<-chan struct{}

	for i := 0; i < 10; i++ {
		slice = append(slice, sig(10*time.Second))
	}
	slice = append(slice, sig(time.Millisecond*1))

	slice = append(slice, sig(12*time.Second))

	start := time.Now()

	<-and(slice...)
	// <-or2(true, slice...)

	fmt.Printf("done after %v", time.Since(start))

}
