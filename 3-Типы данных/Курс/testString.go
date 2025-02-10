package main

import (
	"fmt"
	"strconv"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%d ", sample[i])
	}
	fmt.Printf("\n")
	fmt.Printf("%q\n", sample)

	// s := "hey"
	// rs := []rune([]byte(s)) // cannot convert ([]byte)(s) (type []byte) to type []rune
	// bs := []byte([]rune(s)) // cannot convert ([]rune)(s) (type []rune) to type []byte
	a := string(44)
	fmt.Println(a)
	s := strconv.Itoa(-42)
	i, err := strconv.Atoi("-42")
}
