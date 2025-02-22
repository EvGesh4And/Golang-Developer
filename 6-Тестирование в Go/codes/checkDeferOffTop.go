package main

import "fmt"

func main() {
	fmt.Println(r())
}

func r() (h int) {
	defer func() {
		h = 5
	}()
	h = 3
	return
}
