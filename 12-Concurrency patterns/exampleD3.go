package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	done := make(chan struct{})

	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range 1_000_000_000 {
				s := i * i / (i + 1)
				_ = s
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		start := time.Now()

		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		<-done
		t := time.Since(start)
		fmt.Println("aaaaaaaa", t.Milliseconds())
	}()

	time.Sleep(time.Microsecond)
	close(done)
	wg.Wait()
}
