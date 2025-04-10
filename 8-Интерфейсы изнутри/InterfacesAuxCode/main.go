package main

import "fmt"

type II interface {
	Add() int
}

type T struct {
	a, b int
}

func (t T) Add() int {
	return t.a + t.b
}

func main() {
	var i II
	t1 := T{
		a: 10,
		b: 20,
	}
	i = &t1
	fmt.Println(i.Add())
	fmt.Println(t1.Add())
	t1.a = t1.a * 10
	t1.b = t1.b * 10
	fmt.Println(i.Add())
	fmt.Println(t1.Add())
}
