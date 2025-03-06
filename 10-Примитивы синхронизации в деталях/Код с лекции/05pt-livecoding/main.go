package main

import (
	"log"
	"time"
)

type task struct {
	name  string
	sleep time.Duration
}

func doJob(t task) {
	log.Printf("task %q begin\n", t.name)
	time.Sleep(t.sleep)
	log.Printf("task %q end\n", t.name)
}

func main() {
	tasks := []task{
		{"fast", time.Second / 10},
		{"slow", time.Second * 10},
		{"moderate", time.Second * 2},
	}

	// Сделать так, чтобы задачи выполнялись конкурентно.
	// Дождаться выполнения всех задач.

	log.Println("Начинаем выполнять задачи")
	for _, t := range tasks {
		doJob(t)
	}
	log.Println("ЗАкончили выполнять задачи выполнять задачи")

}
