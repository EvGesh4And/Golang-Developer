package main

import "fmt"

func main() {
	// Общее
	m := map[int]string{5: "vasek"}
	fmt.Printf("%v %#[1]v %[1]T\n", m)

	// Целые числа
	fmt.Println("-----------------")
	fmt.Printf("%b, %[1]d, %[1]o, %[1]x\n", 125)

	str := "Hello, 世界"
	bytes := []byte(str)

	fmt.Println("-----------------")
	fmt.Printf("%s\n", str) // Hello, 世界
	fmt.Printf("%q\n", str) // "Hello, 世界"
	fmt.Printf("%x\n", str) // 48656c6c6f2c20e4b896e7958c

	fmt.Println("-----------------")
	fmt.Printf("%s\n", bytes) // Hello, 世界
	fmt.Printf("%q\n", bytes) // "Hello, 世界"
	fmt.Printf("%x\n", bytes) // 48656c6c6f2c20e4b896e7958c
}
