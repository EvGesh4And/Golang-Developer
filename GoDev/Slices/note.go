package main

import "fmt"

func main() {

	b := [...]int{1, 2, 3, 4, 5, 6, 7}
	a := b[:]
	a = append(a[:2], a[5:]...)
	fmt.Println(b, a)

	b = [...]int{1, 2, 3, 4, 5, 6, 7}
	a = b[:]
	copy(a[:2], a[5:])
	fmt.Println(b, a)
}
