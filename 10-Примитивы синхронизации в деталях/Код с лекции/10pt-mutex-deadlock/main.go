package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu1 := sync.Mutex{}
	mu2 := sync.Mutex{}
	go func() {
		//for {
		fmt.Println(1)
		time.Sleep(time.Second * 10)
		//}
	}()
	//ch := make(chan struct{})
	go func() {

		mu1.Lock()
		time.Sleep(time.Second)
		mu2.Lock()
		// do smth
		mu2.Unlock()
		mu1.Unlock()
		//close(ch)
	}()

	mu2.Lock()
	time.Sleep(time.Second)
	mu1.Lock()
	// do smth
	mu1.Unlock()
	mu2.Unlock()

	//<-ch
}
