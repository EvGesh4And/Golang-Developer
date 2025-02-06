package main

import "fmt"

func fg(int) (b int) {
	b = 1
	a := 2
	return a
}

func main() {
	fmt.Println(fg(4))
}
