package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := gen(nums...)

	ch1 := sq(ch)
	ch2 := sq(ch)

	for num := range merge(ch1, ch2) {
		fmt.Println(num)
	}
}

func merge(ins ...<-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}

	for _, in := range ins {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		defer close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}
