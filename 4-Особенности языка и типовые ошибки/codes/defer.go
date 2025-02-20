package main

import "fmt"

func tt() int {

	res := 3

	defer func() {
		res = 2
	}()

	return res

	// status := 0

	// fmt.Println("first:", status)

	// defer func() {
	// 	fmt.Println("defer-first:", status)
	// }()

	// defer func() func() {
	// 	defer fmt.Println("Tuta")
	// 	fmt.Println("defer-out:", status)
	// 	return func() {
	// 		defer fmt.Println("tuta-tuta")
	// 		fmt.Println("defer-inner:", status)
	// 	}
	// }()()

}

func main() {
	fmt.Println(tt())
}
