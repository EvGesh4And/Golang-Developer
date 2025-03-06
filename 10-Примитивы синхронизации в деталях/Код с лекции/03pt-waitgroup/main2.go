// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	const goCount = 5
	wg := sync.WaitGroup{}
	wg.Add(goCount) // <===
	for i := 0; i < goCount; i++ {
		go func() {
			fmt.Println("go-go-go")
			wg.Done() // <===
		}()
	}
	wg.Wait() // <===
}
