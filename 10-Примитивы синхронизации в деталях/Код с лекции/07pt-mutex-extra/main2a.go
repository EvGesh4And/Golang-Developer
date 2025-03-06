package main

import (
	"fmt"
	"time"
)

func main() {
	const million = 1000000
	//создаем два слайса интов из миллиона элементов каждый
	sl1 := make([]int, million)
	sl2 := make([]int, million)

	//в первом слайсе все элементы делаем равными 1
	//во втором слайсе все элементы делаем равными 2
	for i := 0; i < million; i++ {
		sl1[i] = 1
		sl2[i] = 2
	}
	//Очевидно, что сумма всех элементов обоих слайсов равна три миллиона
	//Но на всякий случай пересчитаем
	var sum int
	//mu := sync.Mutex{}
	go func() {
		for _, v := range sl1 {
			//mu.Lock()
			sum += v
			//mu.Unlock()
		}
	}()

	go func() {
		for _, v := range sl2 {
			//mu.Lock()
			sum += v
			//mu.Unlock()
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(sum)

}
