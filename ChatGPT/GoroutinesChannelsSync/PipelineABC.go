package main

import (
	"context"
	"sync"
)

type Item struct{}
type Result struct{}

// PipelineABC строит конвейер A→B→C, по mX воркеров на стадию.
// Все каналы небуферизированные. При первой ошибке — мягкая остановка всего конвейера.
// Возвращает все успешно полученные Result и список ошибок.
func PipelineABC(
	ctx context.Context,
	in []Item,
	mA, mB, mC int,
	stageA func(context.Context, Item) (Item, error),
	stageB func(context.Context, Item) (Item, error),
	stageC func(context.Context, Item) (Result, error),
) (out []Result, errs []error) {

	// sane defaults
	if mA < 1 {
		mA = 1
	}
	if mB < 1 {
		mB = 1
	}
	if mC < 1 {
		mC = 1
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Каналы между стадиями
	inCh := make(chan Item)   // producer -> A
	out1 := make(chan Item)   // A -> B
	out2 := make(chan Item)   // B -> C
	out3 := make(chan Result) // C -> collector
	errCh := make(chan error) // ошибки всех стадий -> collector

	// Коллекторы запускаем РАНО, чтобы не блокировать писателей
	var wgCollect sync.WaitGroup
	wgCollect.Add(2)

	// Сбор ошибок (только один писатель в errs -> нет гонок)
	go func() {
		defer wgCollect.Done()
		for err := range errCh {
			errs = append(errs, err)
		}
	}()

	// Сбор результатов (только один писатель в out-слайс)
	go func() {
		defer wgCollect.Done()
		for r := range out3 {
			out = append(out, r)
		}
	}()

	// ===== Producer =====
	go func() {
		defer close(inCh)
		for _, it := range in {
			select {
			case <-ctx.Done():
				return
			case inCh <- it:
			}
		}
	}()

	// ===== Stage A =====
	var wgA sync.WaitGroup
	wgA.Add(mA)
	for i := 0; i < mA; i++ {
		go func() {
			defer wgA.Done()
			for it := range inCh {
				next, err := stageA(ctx, it)
				if err != nil {
					// Сначала шлём ошибку, потом гасим весь конвейер и выходим
					errCh <- err
					cancel()
					return
				}
				select {
				case <-ctx.Done():
					return
				case out1 <- next:
				}
			}
		}()
	}
	// Закрыть out1, когда все A-воркеры закончат
	go func() {
		wgA.Wait()
		close(out1)
	}()

	// ===== Stage B =====
	var wgB sync.WaitGroup
	wgB.Add(mB)
	for i := 0; i < mB; i++ {
		go func() {
			defer wgB.Done()
			for it := range out1 {
				next, err := stageB(ctx, it)
				if err != nil {
					errCh <- err
					cancel()
					return
				}
				select {
				case <-ctx.Done():
					return
				case out2 <- next:
				}
			}
		}()
	}
	// Закрыть out2, когда все B-воркеры закончат
	go func() {
		wgB.Wait()
		close(out2)
	}()

	// ===== Stage C =====
	var wgC sync.WaitGroup
	wgC.Add(mC)
	for i := 0; i < mC; i++ {
		go func() {
			defer wgC.Done()
			for it := range out2 {
				res, err := stageC(ctx, it)
				if err != nil {
					errCh <- err
					cancel()
					return
				}
				select {
				case <-ctx.Done():
					return
				case out3 <- res:
				}
			}
		}()
	}

	// Глобальное закрытие out3 и errCh — когда ВСЕ воркеры стадий завершатся
	go func() {
		wgC.Wait()  // C закончили писать в out3
		close(out3) // позволяем коллектору результатов завершиться
		wgA.Wait()  // на случай, если C ушли раньше — дождёмся A/B тоже
		wgB.Wait()
		close(errCh) // после того, как никто больше не сможет писать ошибки
	}()

	// Дождаться, пока коллектора дочитают out3 и errCh и заполнят слайсы
	wgCollect.Wait()
	return out, errs
}
