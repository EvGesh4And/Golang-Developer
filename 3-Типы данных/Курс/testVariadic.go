package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	nums[0] = 13
	fmt.Println(total)
}
func main() {
	sum(5, 7)
	sum(3, 2, 1)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	fmt.Println(nums)
}
