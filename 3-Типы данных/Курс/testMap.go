package main

import "fmt"

func main() {
	var cache map[int]string
	fmt.Printf("%#v\n", cache)
	l := len(cache)   // 0
	v, ok := cache[3] // "", false
	// cache[2] = "ykhj"
	fmt.Println(cache, l, v, ok)
}
