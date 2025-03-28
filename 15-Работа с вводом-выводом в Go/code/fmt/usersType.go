package main

import "fmt"

type myS struct {
	s string
	i int
}

func (m myS) String() string {
	return m.s
}

func main() {

	m := myS{"dee", 1}

	fmt.Println(m)
	fmt.Printf("%v \n", m)
	fmt.Printf("%s \n", m)
}
