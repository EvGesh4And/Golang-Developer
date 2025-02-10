package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	a int8  // 0-й байт
	b int64 // 8-й байт (из-за выравнивания)
	c int8  // 16-й байт
}

func main() {
	fmt.Println(unsafe.Offsetof(Example{}.a)) // 0
	fmt.Println(unsafe.Offsetof(Example{}.b)) // 8
	fmt.Println(unsafe.Offsetof(Example{}.c)) // 16
}
