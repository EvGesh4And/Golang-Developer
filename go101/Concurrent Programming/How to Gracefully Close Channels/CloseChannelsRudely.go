package main

import (
	"fmt"
)

func SafeClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

func main() {
	ch := make(chan int)
	close(ch)
	fmt.Println(SafeClose(ch))
}
