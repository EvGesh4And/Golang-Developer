package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi!")

	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}
}
