package main

import "fmt"

func main() {
	var h chan int
	ch := make(chan int, 1)
	ch <- 2
	select {
	case <-h:
	case a := <-ch:
		fmt.Println(a)
	}
}
