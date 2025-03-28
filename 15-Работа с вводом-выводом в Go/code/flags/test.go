package main

import (
	"flag"
	"fmt"
)

func main() {
	boolPtr := flag.Bool("fork", false, "булев флаг")

	// Парсим аргументы командной строки
	flag.Parse()

	fmt.Println(*boolPtr)
}
