// offTop

package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 5
	f := func(m int) float64 {
		return float64(m + x)
	}

	v := reflect.ValueOf(f)
	fmt.Println("Kind:", v.Kind())       // func
	fmt.Println("IsNil:", v.IsNil())     // false
	fmt.Println("Pointer:", v.Pointer()) // addr of code (not closure)
}
