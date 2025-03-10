package main

import (
	"runtime"
)

func DoSomething() {
	for {
		// do something ...

		runtime.Gosched() // avoid being greedy
	}
}

func main() {
	go DoSomething()
	go DoSomething()
	select {}
}
