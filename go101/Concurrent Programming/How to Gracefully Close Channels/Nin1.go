package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	log.SetFlags(0)

	// ...
	const Max = 100000      // максимальное значение
	const NumSenders = 1000 // количество отправителей

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh — это дополнительный сигнальный канал.
	// Его отправитель — это получатель из dataCh,
	// а его получатели — это все отправители в dataCh.

	// отправители
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// Попытка чтения stopCh, чтобы
				// завершить горутину как можно раньше.
				// В данном примере это не критично.
				select {
				case <-stopCh:
					return
				default:
				}

				// Даже если stopCh закрыт, вторая select
				// может иногда выбирать ветку с отправкой
				// в dataCh, если канал не заблокирован.
				// Это допустимо в этом примере, так что
				// первый select можно было бы и убрать.
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// получатель
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// Получатель из dataCh также является
				// отправителем в stopCh.
				// Закрыть stopCh здесь безопасно.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
