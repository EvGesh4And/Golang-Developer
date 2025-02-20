package main

import "fmt"

func f(x int) {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

var x int

func main() {
	var x = 200
	f(x)
}
