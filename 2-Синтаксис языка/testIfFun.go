package main

import "fmt"

func div(a, b int) (quotient int, ok bool) {
	return 1, true
}

func main() {
	var c int

	if c, ok := div(1, 1); ok {
		fmt.Println(c)
	}
}
