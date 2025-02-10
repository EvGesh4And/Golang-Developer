package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	arr := []int{10, 20, 30}

	ptr := unsafe.Pointer(&arr[0]) // unsafe.Pointer на arr[0]

	// Выводим тип ptr
	fmt.Println("Тип ptr:", reflect.TypeOf((*int)(ptr))) // Тип ptr: *unsafe.Pointer
}
