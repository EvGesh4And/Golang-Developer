package main

func main() {
	i := 3
	switch i {
	case 1:
		var ch chan int
		<-ch // вечная блокировка
	case 2:
		ch := make(chan int)
		close(ch)
		<-ch // всегда ок
	case 3:
		ch := make(chan int)
		<-ch // nomana
	}
}
