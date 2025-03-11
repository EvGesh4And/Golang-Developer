package main

import "fmt"

func SafeSend(ch chan int, value int) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

func main() {
	ch := make(chan int, 1)
	close(ch)
	fmt.Println(SafeSend(ch, 5))
}
