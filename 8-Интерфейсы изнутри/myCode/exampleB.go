package main

import (
	"fmt"
	"reflect"
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

func main() {
	var s Speaker
	h := Human{Greeting: "Hello"}
	s = h
	s.SayHello()

	var ss interface{}
	ss = "ss"
	fmt.Println(reflect.TypeOf(ss).Size())
	fmt.Printf("%v", ss)
}
