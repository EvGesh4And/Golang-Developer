package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	log.SetFlags(0)

	// ...
	const Max = 100000       // максимальное значение
	const NumReceivers = 100 // количество получателей

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// отправитель
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// Единственный отправитель может
				// безопасно закрыть канал в любой момент.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// получатели
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// Читаем значения, пока канал dataCh не будет
			// закрыт и его буферная очередь не опустеет.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
