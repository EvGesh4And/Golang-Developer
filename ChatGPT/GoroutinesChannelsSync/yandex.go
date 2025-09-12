package main

import "sync"

// Есть слайс задач с методом Run() error
// Реализовать функцию func execute(tasks []Task) []error, которая запускает
// каждую задачу в своей горутине и возвращает слайс ошибок.
// В итоговом слайсе не должно быть nil-значений

// Использовать небуферизированный канал

type Task interface {
	Run() error
}

func execute(tasks []Task) []error {

	var res []error

	wg := &sync.WaitGroup{}
	errCh := make(chan error)

	for _, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := task.Run(); err != nil {
				errCh <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		res = append(res, err)
	}

	return res
}
