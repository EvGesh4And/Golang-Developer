package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d) // simulate a workload

			// Changes must be made when
			// cond.L is locked.
			cond.L.Lock()
			values[i] = string('a' + i)

			// Notify when cond.L lock is locked.
			cond.Broadcast()
			cond.L.Unlock()

			// "cond.Broadcast()" can also be put
			// here, when cond.L lock is unlocked.
			//cond.Broadcast()
		}(i)
	}

	// This function must be called when
	// cond.L is locked.
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		// Must be called when cond.L is locked.
		cond.Wait()
	}
}
