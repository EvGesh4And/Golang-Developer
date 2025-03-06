package parallel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// demonstrates goroutine parallel execution
func TestGoroutine1(t *testing.T) {

	buzz := func(str string) {
		for i := 0; i < 9; i++ {
			fmt.Printf("%s %d: buzz\n", str, i)
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("Finish %s\n", str)
	}

	fmt.Println("Start")

	go func() {
		buzz("gorut")
	}()

	buzz("main")

	fmt.Println(runtime.NumGoroutine())
}

// demonstrates parameter passing to goroutine
func TestGoroutine2(t *testing.T) {

	buzz := func(name string) {
		for i := 0; i < 9; i++ {
			fmt.Println("buzz", name)
			time.Sleep(1 * time.Microsecond)
		}
	}

	fmt.Println("Start")

	go buzz("worker") // a simpler way to call

	buzz("main")
}
