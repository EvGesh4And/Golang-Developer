package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 42
	ptr := unsafe.Pointer(&i) // Приведение *int → unsafe.Pointer
	fmt.Println(ptr)
	fptr := (*float64)(ptr) // Приведение unsafe.Pointer → *float64

	fmt.Println(*fptr) // Читаем память как float64 (неожиданные значения!)
}
