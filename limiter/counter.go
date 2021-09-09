package limiter

import (
	"sync/atomic"
	"time"
)

type CounterLimit struct {
	conunter     int64
	limit        int64
	intervalNano int64
	unixNano     int64
}

func NewCounterLimit(interval time.Duration, limit int64) *CounterLimit  {
	return &CounterLimit{
		conunter:     0,
		limit:        limit,
		intervalNano: int64(interval),
		unixNano:     time.Now().Unix(),
	}
}

func (c *CounterLimit) Allow()  bool{
	now := time.Now().Unix()
	if now - c.unixNano > c.intervalNano {
		atomic.StoreInt64(&c.conunter, 0)
		atomic.StoreInt64(&c.unixNano, now)
		return true
	}
	atomic.AddInt64(&c.conunter, 1)
	return c.conunter < c.limit
}
