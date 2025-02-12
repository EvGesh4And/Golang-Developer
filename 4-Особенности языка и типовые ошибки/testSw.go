package main

import "fmt"

func main() {
	// 	switch i := 2; i * 4 {
	// 	case 8:
	// 		j := 2
	// 		fmt.Println(i, j)
	// 	default:
	// 		fmt.Println(i)
	// 	}

	// 	s := make([]int, 2, 3)

	// 	ss := s[0:3]

	// 	fmt.Println(ss)
	// }
	s := "hello"

	defer func(s string) {
		fmt.Println(s)
	}(s)
	defer func() {
		fmt.Println(s)
	}()

	s = "word"
}
