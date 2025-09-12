package main

import (
	"context"
	"math/rand"
	"time"
)

func Retry(ctx context.Context, maxRetries int, base time.Duration, op func(context.Context) error) error {
	if maxRetries < 0 {
		maxRetries = 0
	}
	var err error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err = op(ctx)
		if err == nil {
			return nil
		}
		if attempt == maxRetries {
			break
		}

		// вычислить backoff с защитой от переполнения
		// jitter := [0, backoff)
		d := backoffWithJitter(base, attempt)

		// контекст-aware sleep
		t := time.NewTimer(d)
		select {
		case <-ctx.Done():
			t.Stop()
			return ctx.Err()
		case <-t.C:
		}
	}
	return err // последняя ошибка

}

// backoffWithJitterAdditive возвращает задержку вида:
//
//	base * 2^attempt + U[0, base)
//
// таким образом экспонента сохраняется, а шум добавляется сверху.
func backoffWithJitterAdditive(base time.Duration, attempt int) time.Duration {
	if base <= 0 {
		return 0
	}

	// защита от переполнения при 2^attempt
	if attempt >= 63 {
		attempt = 63
	}
	backoff := base * time.Duration(1<<attempt)

	// ограничим максимум, чтобы не улететь в бесконечность
	const maxBackoff = time.Hour
	if backoff > maxBackoff {
		backoff = maxBackoff
	}

	// добавляем шум в диапазоне [0, base)
	j := rand.Int63n(int64(base))
	return backoff + time.Duration(j)
}
