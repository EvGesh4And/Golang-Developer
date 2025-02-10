package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	newSlice := s
	newSlice[0] = 10
	fmt.Println(s, newSlice)

	s = []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println(len(s), cap(s)) // после первого append
	s = append(s, 6)
	fmt.Println(len(s), cap(s)) // после второго append

	s = make([]int, 2, 4)
	fmt.Println(len(s), cap(s)) // изначально
	s = append(s, 3)
	fmt.Println(len(s), cap(s)) // после первого append
	s = append(s, 4, 5)
	fmt.Println(len(s), cap(s)) // после второго append
	s = append(s, 6)
	fmt.Println(len(s), cap(s)) // после третьего append

	s = []int{1, 2, 3}
	s = s[:2]              // обрезаем срез
	s = append(s, 4, 5, 6) // пытаемся добавить 3 элемента
	fmt.Println(s)

	s = []int{1, 2, 3}
	t := make([]int, len(s)-1)
	copy(t, s)
	t[0] = 100
	fmt.Println(s, t)

	s = []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s, len(s), cap(s)) // после первого append
	s = s[:2]
	fmt.Println(s, len(s), cap(s)) // после обрезки

	s = make([]int, 2, 5)

	func(ss []int) {
		ss = append(ss, 6)
	}(s)

	fmt.Println(s[2:5])
}
