package ratelimiter

import (
	"sync/atomic"
)

// Limiter contains data used uinternally
type Limiter struct {
	counter uint64
}

func New(maximumValue uint64, initialValue uint64) *Limiter {
	return &Limiter{
		counter: initialValue,
	}
}

func (l *Limiter) Take(count uint64) {
	atomic.AddUint64(&l.counter, count)
}
