package main

import "fmt"

func main() {

	a := []int{1, 2, 3}

	a = append(a, 3, 4, 6, 7, 8, 9)

	fmt.Println(cap(a))
}
