package main

import (
	"fmt"
)

func SafeClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// Возврат можно изменить в функции defer.
			justClosed = false
		}
	}()

	// предполагаем, что ch != nil здесь.
	close(ch)   // вызовет panic, если канал уже закрыт
	return true // то же самое, что justClosed = true; return
}

func main() {
	ch := make(chan int)
	close(ch)
	fmt.Println(SafeClose(ch))
}
