package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = [2]int{42, 99}
	ptr := unsafe.Pointer(&arr[0])                             // Получаем указатель на arr[0]
	ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0])) // Смещаемся на 8 байт (int64)

	fmt.Println(*(*int)(ptr)) // Читаем arr[1] (99)

	myArr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	mySlice := myArr[3:5:6]

	ptr = unsafe.Pointer(&mySlice[0])

	ptr = unsafe.Pointer(uintptr(ptr) - 3*unsafe.Sizeof(arr[0]))

	fmt.Println(*(*int)(ptr))
}
