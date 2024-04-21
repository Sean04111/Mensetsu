package util

import (
	"context"
	"sync/atomic"
	"time"
)

type Limiter struct {
	stone     atomic.Int64
	threshold int64
}

func NewLimiter(threshold int64, ctx context.Context) *Limiter {
	l := &Limiter{}
	l.threshold = threshold
	go l.Produce(ctx)
	return l
}

func (l *Limiter) Consume(n int64) (consumed int64) {
	if l.stone.Load() <= n {
		val := l.stone.Load()
		for !l.stone.CompareAndSwap(val, 0) {
			val = l.stone.Load()
		}
		return val
	} else {
		val := l.stone.Load()
		for !l.stone.CompareAndSwap(val, val-n) {
			val = l.stone.Load()
		}
		return n
	}

}

func (l *Limiter) Produce(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		value := l.stone.Load()
		for !l.stone.CompareAndSwap(value, value+l.threshold) {
			value = l.stone.Load()
		}

		time.Sleep(time.Second)
	}
}
