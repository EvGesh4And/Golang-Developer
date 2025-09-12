package main

import (
	"errors"
	"fmt"
	"sync"
)

func ParallelMap(in []int, workers int, work func(int) (int, error)) (outs []int, errs []error) {
	if workers <= 0 {
		workers = 1
	}
	outs = make([]int, len(in))
	errs = make([]error, len(in))

	type input struct {
		idx int
		val int
	}
	inputsCh := make(chan input) // небуферизированный

	// продюсер
	go func() {
		for i, v := range in {
			inputsCh <- input{i, v}
		}
		close(inputsCh)
	}()

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		var in input
		wg.Add(1)
		go func() {
			defer wg.Done()
			// опционально: защита от паник
			defer func() {
				if r := recover(); r != nil {
					errs[in.idx] = errors.New(fmt.Sprintf("panic at idx=%d, val=%d: %v\n", in.idx, in.val, r))
				}
			}()

			for in = range inputsCh {
				out, err := work(in.val)
				outs[in.idx] = out
				errs[in.idx] = err // 1:1 модель — nil допустим
			}
		}()
	}
	wg.Wait()
	return
}
