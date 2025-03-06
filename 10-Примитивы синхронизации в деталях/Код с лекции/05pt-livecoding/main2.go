package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type task struct {
	name  string
	sleep time.Duration
}

func doJob(t task) {
	fmt.Printf("task %q begin\n", t.name)
	time.Sleep(t.sleep)
	fmt.Printf("task %q end\n", t.name)
}

func main() {
	var wg sync.WaitGroup

	tasks := []task{
		{"fast", time.Second / 10},
		{"slow", time.Second * 10},
		{"moderate", time.Second * 2},
	}
	log.Println("Начинаем выполнять задачи")
	// Сделать так, чтобы задачи выполнялись конкурентно.
	// Дождаться выполнения всех задач.
	for _, t := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doJob(t)
		}()
	}

	wg.Wait()
	log.Println("ЗАкончили выполнять задачи выполнять задачи")
}
