package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Общие параметры
	capacity := 5
	tasks := make([]int, 0, capacity)

	mu := sync.Mutex{}     // Мьютекс для управления задачами
	c := sync.NewCond(&mu) // sync.Cond для ожидания и пробуждения

	// Производитель
	producer := func(id int) {
		for i := 0; i < 20; i++ { // Ограничиваем количество итераций
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

			mu.Lock()
			// Ожидаем, пока не освободится место в очереди
			for len(tasks) == capacity {
				c.Wait()
			}

			task := rand.Intn(100)
			tasks = append(tasks, task)
			log.Printf("[Producer %d] Добавил задачу %d", id, task)

			// Разбудить одного потребителя
			c.Signal()

			mu.Unlock()
		}
	}

	// Потребитель
	consumer := func(id int) {
		for i := 0; i < 20; i++ { // Ограничиваем количество итераций
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(700)))

			mu.Lock()
			// Ожидаем, пока не появится задача в очереди
			for len(tasks) == 0 {
				c.Wait()
			}

			task := tasks[len(tasks)-1]
			tasks = tasks[:len(tasks)-1] // Забираем задачу
			log.Printf("[Consumer %d] Обработал задачу %d", id, task)

			// Разбудить одного производителя
			c.Signal()

			mu.Unlock()
		}
	}

	// Запускаем производителей
	numProducers := 5
	for i := 0; i < numProducers; i++ {
		go producer(i + 1)
	}

	// Запускаем потребителей
	numConsumers := 5
	for i := 0; i < numConsumers; i++ {
		go consumer(i + 1)
	}

	// Ждём завершения всех горутин
	time.Sleep(time.Second * 10)
	log.Println("Все задачи обработаны!")
}
