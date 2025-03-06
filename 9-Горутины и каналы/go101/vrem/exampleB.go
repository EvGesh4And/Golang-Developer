package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetingsMod(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	// Notify a task is finished.
	wg.Done() // <=> wg.Add(-1)
}

func main() {
	log.SetFlags(0)
	wg.Add(2) // register two tasks.
	go SayGreetingsMod("hi!", 10)
	go SayGreetingsMod("hello!", 10)
	wg.Wait() // block until all tasks are finished.
}
