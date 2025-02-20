package main

import "fmt"

func main() {
	classicIf()
	withNeyavnyBlockIf()
}

func classicIf() {
	if x := 10; x > 5 { // x создаётся в неявном блоке
		fmt.Println("x больше 5")
	} // Здесь x уничтожается
}

func withNeyavnyBlockIf() {
	{ // Неявный блок if
		x := 10    // x создаётся
		if x > 5 { // Используется x
			fmt.Println("x больше 5")
		}
	} // Здесь x уничтожается
}
