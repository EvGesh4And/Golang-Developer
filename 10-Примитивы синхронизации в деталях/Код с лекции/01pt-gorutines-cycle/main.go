package main

import (
	"fmt"
	"time"
)

func main() {
	//start := time.Now()
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("go-go-go")
		}()
	}
	//fmt.Println(time.Since(start).String())
	time.Sleep(time.Second)
}
