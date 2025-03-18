package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]bool)
	wg := sync.WaitGroup{}

	const enableLogging = true
	//try to change this to false ;-)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				current := i*10 + j
				if enableLogging {
					fmt.Printf("Собираемся записывать значение: %d \n", current)
				}

				m[current] = true
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Мапа содержит %d элементов \n", len(m))
	fmt.Println("Чудом не упали")
}
