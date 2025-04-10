package main

import "fmt"

func main() {
	var i interface{} = "3.14"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	r, ok := i.(fmt.Stringer)
	fmt.Println(r, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}
