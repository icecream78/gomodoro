package pomodoro

import (
	"sync"
)

type Timer struct {
	seconds int
	current int
	border  int
	mx      sync.Mutex
}

func calculateNotifyTime(stageTime int, percent int) int {
	perc := (float32(percent) / 100)
	return int(float32(stageTime) * perc)
}

func NewTimer(t int, border int) *Timer {
	return &Timer{
		seconds: t,
		current: t,
		border:  calculateNotifyTime(t, border),
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

func (t *Timer) NeedNotify() bool {
	t.mx.Lock()
	now := t.current
	t.mx.Unlock()
	return now <= t.border
}
