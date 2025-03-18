package main

import "sync"

func and(chaneles ...<-chan struct{}) <-chan struct{} {
	andDone := make(chan struct{})
	go func() {
		defer close(andDone)
		wg := sync.WaitGroup{}
		for _, ch := range chaneles {
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <-ch:
				case <-andDone:
				}
			}()
		}
		wg.Wait()
	}()
	return andDone
}
