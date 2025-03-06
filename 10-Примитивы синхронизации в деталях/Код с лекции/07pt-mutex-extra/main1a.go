package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	mu := sync.Mutex{}
	go stih1(&mu)
	go stih2(&mu)
	time.Sleep(time.Second)
}
func stih1(mu *sync.Mutex) {

	fmt.Println("Три мудреца в одном тазу")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Пустились по морю в грозу")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Будь попрочнее старый таз,")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Длиннее был бы мой рассказ")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("---------------------------------")

}

func stih2(mu *sync.Mutex) {

	fmt.Println("Шалтай-болтай сидел на стене")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Шалтай-болтай свалился во сне")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("И вся королевская конница, вся королевская рать")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Не может Шалтая-Болтая собрать")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("---------------------------------")

}
