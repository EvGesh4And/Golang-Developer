package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source() <-chan int32 {
	// c must be a buffered channel.
	c := make(chan int32, 1)
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3)+1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}

func main() {
	var rnd int32
	// Blocking here until one source responses.
	select {
		select {
	case rnd = <-source():
	case rnd = <-source():
	case rnd = <-source():
	}
	fmt.Println(rnd)
}
