package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, 0666)
	fmt.Println(f, err)

	// Через реализацию интерфейса io.Reader
	buf := make([]byte, 10)
	n, err := f.Read(buf)

	fmt.Println(n, err)
	fmt.Println(buf, n, err)
	f.Close()

	f, err = os.OpenFile("data.txt", os.O_RDONLY, 0666)
	fmt.Println(f, err)

	n, err = io.ReadFull(f, buf)
	fmt.Println(n, err)

	f, err = os.OpenFile("data.txt", os.O_RDONLY, 0666)
	fmt.Println(f, err)

	// Через функцию ReadAll
	b, err := io.ReadAll(f)
	fmt.Println(b)
}
