package main

import "fmt"

func main() {
	classicFor()
	withNeyavnyBlockFor()
}

func classicFor() {
	for i := 0; i < 3; i++ { // i создаётся в неявном блоке
		x := i * 2
		fmt.Println(x)
	} // Здесь i уничтожается
}

func withNeyavnyBlockFor() {
	{ // Неявный блок for
		i := 0 // i создаётся в неявном блоке
	tuta: // метка для goto
		if i < 3 {
			// начало тела функции
			x := i * 2
			fmt.Println(x)
			// конец тела функции
			i++
			goto tuta // возвращаемся на метку
		}
	} // Здесь i уничтожается
}
