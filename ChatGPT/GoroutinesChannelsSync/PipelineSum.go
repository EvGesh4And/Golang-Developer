package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(PipelineSum(ctx, 999_000_000))
}

func PipelineSum(ctx context.Context, N int) (int, error) {
	res := sum(square(ctx, gen(ctx, N)))
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	return res, nil
}

func gen(ctx context.Context, N int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 1; i <= N; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()

	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for num := range in {
			select {
			case <-ctx.Done():
				return
			case out <- num * num:
			}
		}
	}()

	return out
}

func sum(in <-chan int) (res int) {
	for num := range in {
		res += num
	}
	return
}
