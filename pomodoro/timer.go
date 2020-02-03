package pomodoro

import (
	"sync"
)

type Timer struct {
	seconds int
	current int
	mx      sync.Mutex
}

func NewTimer(t int) *Timer {
	return &Timer{
		seconds: t,
		current: t,
	}
}

// Ticker implementation
func (t *Timer) Tick() {
	t.Decrement()
}

func (t *Timer) Refresh() {
	t.Reset()
}

func (t *Timer) Finished() bool {
	t.mx.Lock()
	now := t.current
	t.mx.Unlock()
	return now == 0
}

func (t *Timer) State() int {
	t.mx.Lock()
	now := t.current
	t.mx.Unlock()
	return now
}

func (t *Timer) Minus(sec int) *Timer {
	t.mx.Lock()
	t.current -= sec
	t.mx.Unlock()
	return t
}

func (t *Timer) Decrement() *Timer {
	return t.Minus(1)
}

func (t *Timer) Reset() *Timer {
	t.mx.Lock()
	t.current = t.seconds
	t.mx.Unlock()
	return t
}
