package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	const N = 5
	var values [N]int32

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			values[i] = 50 + rand.Int31n(50)
			log.Println("Done:", i)
			wg.Done() // <=> wg.Add(-1)
		}()
	}

	wg.Wait()
	// All elements are guaranteed to be
	// initialized now.
	log.Println("values:", values)
}
