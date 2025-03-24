package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("read.txt")
	fmt.Println(err)
	fout, err := os.OpenFile("write.txt", os.O_WRONLY, 0666)

	_, err = io.Copy(fout, f)
}
