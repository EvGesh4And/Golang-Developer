package main

import (
	"fmt"
)

func Concat(sls ...[]int) []int {
	sumLen := 0

	for _, v := range sls {
		sumLen += len(v)
	}

	resSlice := make([]int, sumLen)

	currPos := 0
	for _, v := range sls {
		copy(resSlice[currPos:], v)
		currPos += len(v)
	}

	return resSlice
}

func main() {
	res := Concat([]int{1, 3, 5}, []int{2, 5}, []int{}, []int{6, 8}, []int{7})

	fmt.Println(res)
}
