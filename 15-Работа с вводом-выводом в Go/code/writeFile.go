package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("write.txt")
	fmt.Println(err)

	b, err := os.ReadFile("read.txt")
	fmt.Println(err)

	f.Write(b)
}
