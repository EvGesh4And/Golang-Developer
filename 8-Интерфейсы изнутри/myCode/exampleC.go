package main

import "fmt"

type Animal interface {
	Name() string
	Age() int
}

type Cat string

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) Age() int {
	return 10
}

type Homa string

func (h Homa) Name() string {
	return string(h)
}

func (h Homa) Age() int {
	return 10
}

type Mouse string

func (m Mouse) Name() string {
	return string(m)
}

func (m Mouse) Age() int {
	return 10
}

func (m Mouse) LenHv() int {
	return 200
}

func main() {
	var i Animal

	i = Cat("a")

	switch a := i.(type) {
	case Cat, Homa:
		fmt.Println(a.Name())
	case Mouse:
		fmt.Println(a.LenHv())
	}
}
