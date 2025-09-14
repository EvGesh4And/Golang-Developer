package main

import (
	"context"
)

func MergeSorted(ctx context.Context, ins ...<-chan int) <-chan int {
	out := make(chan int)

	// edge-case: нет входов
	if len(ins) == 0 {
		close(out)
		return out
	}

	go func() {
		defer close(out)

		type item struct {
			val int
			ok  bool
		}
		heads := make([]item, len(ins))

		// начальное чтение «голов»
		for i, ch := range ins {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				heads[i] = item{v, ok}
			}
		}

		for {
			// найти минимальную «голову»
			minIdx := -1
			for i, it := range heads {
				if !it.ok {
					continue
				}
				if minIdx == -1 || it.val < heads[minIdx].val {
					minIdx = i
				}
			}
			if minIdx == -1 {
				return // все входы пусты
			}

			// отдать минимальный
			select {
			case <-ctx.Done():
				return
			case out <- heads[minIdx].val:
			}

			// дочитать следующую «голову» из того же входа
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ins[minIdx]:
				heads[minIdx] = item{v, ok}
			}
		}
	}()

	return out
}
