package main

import "fmt"

func main() {
	t := make([]int, 13, 13)

	fmt.Println(t, len(t), cap(t))

	t = append(t, 2)
	fmt.Println(t, len(t), cap(t))

	a := []int{1, 2, 3}
	b := a[1:2:2]
	fmt.Println(b == nil, b, len(b), cap(b))
	for _, x := range a {
		if f(x) {
			b = append(b, x*100)
		}
	}

	fmt.Println(a)
	aa := []*int{&a[0], &a[1]}
	clear(aa[1:])
	fmt.Println(aa)
}

func f(i int) bool {
	return i < 10
}
