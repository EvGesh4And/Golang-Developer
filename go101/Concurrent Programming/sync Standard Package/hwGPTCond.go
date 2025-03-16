package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type MyCounter struct {
	m sync.Mutex
	v int
}

func (c *MyCounter) Add() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v++
}

func (c *MyCounter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.v
}

func main() {
	count := MyCounter{}

	c := sync.NewCond(&sync.Mutex{})

	capacity := 5
	tasks := make([]int, 0, capacity)

	producer := func() {
		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		c.L.Lock()
		defer c.L.Unlock()

		for len(tasks) == 5 {
			c.Wait()
		}
		tasks = append(tasks, rand.Intn(10))
		c.Broadcast()
		count.Add()
	}

	for range 100 {
		go producer()
	}

	consumer := func() {
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		c.L.Lock()
		defer c.L.Unlock()

		for len(tasks) == 0 {
			c.Wait()
		}
		log.Println(tasks[len(tasks)-1])
		tasks = tasks[0 : len(tasks)-1]
		c.Broadcast()
		log.Println("val:", count.Value())
	}

	for range 100 {
		go consumer()
	}

	time.Sleep(time.Second * 6)
}
