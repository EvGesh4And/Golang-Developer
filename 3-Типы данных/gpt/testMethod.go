package main

import "fmt"

func (p Person) Speak() {
	fmt.Println("Hi, I'm", p.Name)
}

func main() {
	type Person struct {
		Name string
	}
}
