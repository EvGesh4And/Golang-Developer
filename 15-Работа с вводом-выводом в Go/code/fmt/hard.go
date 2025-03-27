package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	sl := []int{1, 2, 3}
	mp := map[string]string{"qwe": "asd"}
	ch := make(chan int, 3)

	var st struct {
		s string
		i int
	}

	fmt.Printf("%p\n%v\n%#v\n\n", &arr, arr, arr)
	fmt.Printf("%p\n%v\n%#v\n\n", sl, sl, sl)
	fmt.Printf("%p\n%v\n%#v\n\n", mp, mp, mp)
	fmt.Printf("%p\n%v\n%#v\n\n", ch, ch, ch)
	fmt.Printf("%v\n%#v\n\n", st, st)
}
