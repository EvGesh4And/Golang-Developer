package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

		time.Sleep(3 * time.Second)
		wg.Done()
	}()

	go func() {
		fmt.Println("Я подожду")
		wg.Wait()
		fmt.Println("Я дождался")

	}()

	go func() {
		fmt.Println("Я тоже подожду")
		wg.Wait()
		fmt.Println("Я тоже дождался")
	}()

	time.Sleep(5 * time.Second)

}
