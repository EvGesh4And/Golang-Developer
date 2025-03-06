package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	v := 0
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			outV := 123 // сетевой вызов на 5 секунд
			var err error
			mu.Lock() // <===
			v, err = outV, errors.New("1")
			if err != nil {
				return
			}
			mu.Unlock()

			fmt.Println(outV)
		}()
	}
	wg.Wait()
	fmt.Println(v)
}
