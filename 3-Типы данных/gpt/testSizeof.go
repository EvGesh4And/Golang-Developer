package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {

	var x int64
	fmt.Println(unsafe.Sizeof(x)) // 8 байт

	var y struct {
		a int8
		b int64
		c int8
	}
	fmt.Println(unsafe.Sizeof(y)) // 24 байта (из-за выравнивания)

	var com complex128

	fmt.Println(unsafe.Sizeof(com))

	var z struct {
		a int8
		b complex128
		c int8
	}

	fmt.Println(unsafe.Alignof(z))

	type S2 struct {
		m sync.Mutex // Может требовать 8, 16 или 64 байта в зависимости от платформы
	}

	fmt.Println("Alignof(sync.Mutex):", unsafe.Alignof(S2{}))

	type MyInt int

	fmt.Println(MyInt(5))

	const (
		A = iota
		B
	)

	fmt.Println(A, B)
}
