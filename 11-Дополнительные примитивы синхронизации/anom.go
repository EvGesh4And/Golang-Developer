package main

import (
	"fmt"
	"time"
)

func main() {
	text := ""
	isInit := false

	go func() {
		text = "go-go-go"
		isInit = true
	}()

	for !isInit {
		time.Sleep(time.Microsecond)
	}

	fmt.Println(text)
}