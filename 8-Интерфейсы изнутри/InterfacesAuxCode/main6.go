package main

import "fmt"

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	sl := make([]any, 0, len(names))
	for _, v := range names {
		sl = append(sl, v)
	}
	PrintAll(sl)
}
