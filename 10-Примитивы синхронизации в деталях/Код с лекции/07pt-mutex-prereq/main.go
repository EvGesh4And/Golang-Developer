package main

import (
	"fmt"
	"sync"
)

func main() {

	//fmt.Println(runtime.GOMAXPROCS(8))

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	v := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()

			v++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v)
}
