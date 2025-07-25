package main

import "fmt"

type Aboutable interface {
	About() string
}

type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book: " + book.name
}

func main() {
	// Упаковка *Book в интерфейс Aboutable
	var a Aboutable = &Book{"Go 101"}
	fmt.Println(a) // &{Go 101}

	// i — пустой интерфейс
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable реализует interface{}, можно присвоить
	i = a
	fmt.Println(i) // &{Go 101}
}
