package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(10)

	wg.Add(-10)

	wg.Wait()
	fmt.Println("daa")
}
