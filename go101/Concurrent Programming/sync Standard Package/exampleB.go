package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {

	const N = 5
	var values [N]int32

	var wgA, wgB sync.WaitGroup
	wgA.Add(N)
	wgB.Add(1)

	for i := 0; i < N; i++ {
		go func() {
			wgB.Wait() // wait a notification
			log.Printf("values[%v]=%v \n", i, values[i])
			wgA.Done()
		}()
	}

	// The loop is guaranteed to finish before
	// any above wg.Wait calls returns.
	for i := 0; i < N; i++ {
		values[i] = 50 + rand.Int31n(50)
	}
	// Make a broadcast notification.
	wgB.Done()
	wgA.Wait()
}
