package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Create("buf.txt")

	bw := bufio.NewWriter(f)

	bw.Write([]byte("sda"))

	bw.Flush()
}
