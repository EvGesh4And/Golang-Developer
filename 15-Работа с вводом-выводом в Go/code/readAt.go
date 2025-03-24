package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("read.txt")
	b := make([]byte, 50)
	f.ReadAt(b, 1)
	fmt.Println(string(b), err)
}
