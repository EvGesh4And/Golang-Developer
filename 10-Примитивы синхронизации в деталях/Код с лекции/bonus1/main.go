package main

import (
	"log"
	"sync"
)

func main() {

	once := sync.Once{}

	a := 100
	once.Do(func() { ofunc(a) })
	once = sync.Once{}
	once.Do(func() { ofunc(a) })
	once.Do(func() { ofunc(a) })

}

func ofunc(i int) {
	log.Println(i)
}
