package main

import (
	"context"
)

func MergeSorted(ctx context.Context, ins ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		type item struct {
			val int
			ok  bool
		}
		heads := make([]item, len(ins))

		// начальное чтение
		for i, ch := range ins {
			v, ok := <-ch
			heads[i] = item{v, ok}
		}

		for {
			// выбираем минимальный из heads
			minIdx := -1
			for i, it := range heads {
				if !it.ok {
					continue
				}
				if minIdx == -1 || it.val < heads[minIdx].val {
					minIdx = i
				}
			}
			if minIdx == -1 { // все каналы пусты
				return
			}

			select {
			case <-ctx.Done():
				return
			case out <- heads[minIdx].val:
			}

			// читаем следующий из выбранного канала
			v, ok := <-ins[minIdx]
			heads[minIdx] = item{v, ok}
		}
	}()

	return out
}
