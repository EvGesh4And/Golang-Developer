package main

import (
	"log"
	"sync"
)

func main() {

	o := sync.Once{}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		o.Do(func() {
			log.Println("1")
		})
	}()

	go func() {
		defer wg.Done()
		o.Do(func() {
			log.Println("2")
		})
	}()

	wg.Wait()
}
