package main

import (
	"fmt"
)

type MyVeryOwnStringer struct{ s string }

func (s MyVeryOwnStringer) String() string {
	return "my string representation of MyVeryOwnStringer"
}

func main() {
	// my string representation of MyVeryOwnStringer{}
	fmt.Println(MyVeryOwnStringer{"hello"})
}
