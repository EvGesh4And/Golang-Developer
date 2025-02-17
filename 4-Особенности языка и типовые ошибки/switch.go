package main

import "fmt"

func main() {
	classicSwitch()
	withNeyavnyBlockSwitch()
}

func classicSwitch() {
	switch x := 2; x { // x создаётся в неявном блоке
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	} // Здесь x уничтожается
}

func withNeyavnyBlockSwitch() {
	{ // Неявный блок for
		x := 2 // x создаётся в неявном блоке
		if x == 1 {
			fmt.Println("Один")
		}
		if x == 2 {
			fmt.Println("Два")
		}
	} // Здесь x уничтожается
}
