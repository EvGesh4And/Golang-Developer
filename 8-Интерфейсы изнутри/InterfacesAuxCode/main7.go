package main

import (
	"fmt"
	"unsafe"
)

type Speaker interface {
	SayHello()
}

type Human struct {
	Greeting string
}

func (h Human) SayHello() {
	fmt.Println(h.Greeting)
}

func meet(s Speaker) {
	s.SayHello()
}

func main() {
	var s Speaker
	fmt.Println(unsafe.Sizeof(s))
	h := Human{Greeting: "Hello"}
	s = Speaker(h)
	s.SayHello()

	meet(Human{Greeting: "Hello :)"})
}
