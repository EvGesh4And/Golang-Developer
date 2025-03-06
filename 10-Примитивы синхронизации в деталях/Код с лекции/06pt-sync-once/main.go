package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	var wg sync.WaitGroup
	for i := 9; i > 0; i-- {
		wg.Add(1)
		go func(i int) {
			once.Do(
				func() {
					Printer(fmt.Sprintf("%d", i))
				},
			)
			wg.Done()

		}(i)
	}
	wg.Wait()
}

func Printer(title string) {
	fmt.Println(title)
}
