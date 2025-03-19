package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("exit normally.")
	}()
	fmt.Println("hi!")
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	panic("bye!")
	fmt.Println("unreachable")
}
