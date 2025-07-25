package main

import "fmt"

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"
	// Сейчас динамический тип y — "string".
	var w Writer
	var ok bool

	// Тип DummyWriter реализует и Writer, и interface{}.
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// Динамический тип y — "string", он не реализует Writer.
	w, ok = y.(Writer)
	fmt.Println(w, ok) // <nil> false
	w = y.(Writer)     // приведёт к панике
}
