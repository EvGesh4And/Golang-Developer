package main

import (
	"fmt"
)

type Hound interface {
	bark(number int)
}

type Retriever interface {
	Hound
	bark(int) // duplicate method
}

func main() {
	fmt.Println("hello")
}

type dog struct {
}

func (d dog) bark(count int) {

}
