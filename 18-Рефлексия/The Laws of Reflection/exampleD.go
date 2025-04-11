package main

import (
	"fmt"
	"reflect"
)

func main() {
	type st struct {
		I int
		f float64
	}

	t := st{10, 20}

	s := reflect.ValueOf(t)

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v \n", i, typeOfT.Field(i).Name, f.Type(), f)
	}
}
