package main

import (
	"context"
	"sync"
)

func FetchAll(ctx context.Context, urls []string, k int, fetch func(context.Context, string) error) []error {
	if k < 1 {
		k = 1
	}

	var errs []error

	type input struct {
		idx int
		url string
	}
	inputCh := make(chan input)
	errCh := make(chan error)

	go func() {
		defer close(inputCh)
		for i, url := range urls {
			select {
				case <-ctx.Done():
					return
				case inputCh <- input{i, url}:
				}
			}
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for in := range inputCh {
				if err := fetch(ctx, in.url); err != nil {
					errCh <- err
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		errs = append(errs, err)
	}

	return errs
}
