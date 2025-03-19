package main

func main() {

	var aChannel chan int

	for v := range aChannel {
		// use v
	}
}
