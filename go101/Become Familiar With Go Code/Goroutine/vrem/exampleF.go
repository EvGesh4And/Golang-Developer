package main

import "fmt"

func main() {
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i+x)
		}
		x = 10
	}()
	fmt.Println()
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i+x)
			}()
		}
		x = 10
	}()
}
