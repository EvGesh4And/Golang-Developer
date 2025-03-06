package main

import "fmt"

type emptyStruct struct{}

func main() {
	const goCount = 5

	ch := make(chan struct{})
	for i := 0; i < goCount; i++ {
		go func(i int) {
			fmt.Printf("go-go-go %d\n", i)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < goCount; i++ {
		<-ch
	}
}
