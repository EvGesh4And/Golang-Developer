package main

import "fmt"

func deleteElementLeak(slice []*int, index int) []*int {
	copy(slice[index:], slice[index+1:]) // Сдвигаем элементы влево
	return slice[:len(slice)-1]          // Обрезаем срез
}

func main() {
	a, b, c, d := 1, 2, 3, 4
	slice := []*int{&a, &b, &c, &d}

	fmt.Println("Before delete:", slice)

	slice = deleteElementLeak(slice, 1) // Удаляем второй элемент (b)

	fmt.Println("After delete:", slice)

	fmt.Println("After delete:", slice[:cap(slice)])
}
