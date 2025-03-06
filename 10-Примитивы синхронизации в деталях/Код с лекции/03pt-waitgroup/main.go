package main

import (
	"fmt"
	"sync"
)

func main() {
	const goCount = 5
	wg := sync.WaitGroup{}
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func(wg1 *sync.WaitGroup) {
			fmt.Println("go-go-go")
			wg1.Done() // <===
		}(&wg)
	}
	wg.Wait() // <===
	fmt.Println("DONE")
}
