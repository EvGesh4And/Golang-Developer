package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("kind is float64:", reflect.ValueOf(x).Kind())
	fmt.Println("value:", reflect.ValueOf(x).Float())

	type metr int
	var y metr = metr(4)
	fmt.Println("type:", reflect.TypeOf(y))
	fmt.Println("value:", reflect.ValueOf(y))
	fmt.Println("value:", reflect.ValueOf(y).String())
	fmt.Println("kind is float64:", reflect.ValueOf(y).Kind())
	fmt.Println("value:", reflect.ValueOf(y).Int())

	var z string = "zxc"
	fmt.Println("type:", reflect.TypeOf(z))
	fmt.Println("value:", reflect.ValueOf(z))
	fmt.Println("value:", reflect.ValueOf(z).String())
	fmt.Println("kind is float64:", reflect.ValueOf(z).Kind())
	fmt.Println("value:", reflect.ValueOf(z).String())
}
