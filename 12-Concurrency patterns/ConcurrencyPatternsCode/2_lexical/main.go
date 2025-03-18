package main

import (
	"fmt"
	"sync"
)

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int)

		n := 10
		var wg sync.WaitGroup
		wg.Add(n)

		for k := 1; k <= n; k++ {
			k := k

			go func() {
				defer wg.Done()

				for i := 0; i <= 5; i++ {
					results <- i + k*n
				}
			}()
		}

		go func() {
			wg.Wait()
			defer close(results)
		}()

		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()

	consumer(results)
}
