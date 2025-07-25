package main

import "fmt"

func main() {
	// Компилятор выводит тип 123 как int.
	var x interface{} = 123

	// Случай 1:
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// Случай 2:
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// Случай 3:
	a = x.(float64) // приведёт к панике
}
