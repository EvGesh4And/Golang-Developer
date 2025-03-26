package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{'a', 'b', 'c'})
	b := []byte{1}
	buf.Read(b)
	contents := buf.String()

	fmt.Println(contents)
}
