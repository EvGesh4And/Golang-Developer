package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	rwm := sync.RWMutex{}

	a := 'a'

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for a < 'z' {
				rwm.RLock()
				log.Println(a)
				rwm.RUnlock()
			}
		}()
	}

	for range 37 {
		func() {
			defer rwm.Unlock()
			rwm.Lock()
			a++
		}()
	}

	wg.Wait()
	log.Println(a)
}
