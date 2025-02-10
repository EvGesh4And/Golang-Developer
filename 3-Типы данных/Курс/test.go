package main

import "fmt"

func main() {
	var s []int

	s = make([]int, 0, 5)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)

	fmt.Println(len(s), cap(s))

	s = append(s, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25)

	fmt.Println(len(s), cap(s))
}
