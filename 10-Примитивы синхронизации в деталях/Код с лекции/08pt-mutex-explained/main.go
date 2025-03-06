package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	v := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// BEGIN ()
			old_v := v         // 1         | //1
			new_v := old_v + 1 //2  | //2
			v = new_v          // 2          | //2
			//COMMIT()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v)
}
