package main

func main() {
	i := 3
	switch i {
	case 1:
		var ch chan int
		close(ch) // panic
	case 2:
		ch := make(chan int)
		close(ch)
		close(ch) // panic
	case 3:
		ch := make(chan int)
		close(ch) // nomana
	}

}
