package pomodoro

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
)

type colorFunc func(a ...interface{}) string

var (
	progressColor  = color.New(color.FgRed, color.Bold).SprintFunc()
	finishingColor = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	finishedColor  = color.New(color.FgGreen, color.Bold).SprintFunc()
)

func PadLeft(original string, padSymb string, length int) string {
	if length <= 0 || len(original) >= length {
		return original
	}
	s := original
	for i := 0; i < length-len(original); i++ {
		s = padSymb + s
	}
	return s
}

type Timer struct {
	seconds      int
	current      int
	notifyBorder int
	mx           sync.Mutex
}

func NewTimer(t int, borderSec int) *Timer {
	return &Timer{
		seconds:      t,
		current:      t,
		notifyBorder: borderSec,
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

func (t *Timer) getMinutesSeconds() (min int, sec int) {
	t.mx.Lock()
	now := t.current
	t.mx.Unlock()
	if now > 0 {
		sec = now % 60
		min = (now - sec) / 60
	}
	return
}

func (t *Timer) isEnding() bool {
	t.mx.Lock()
	now := t.current
	t.mx.Unlock()
	return now <= t.notifyBorder
}

func (t *Timer) getColorFunc() colorFunc {
	if t.Finished() {
		return finishedColor
	}
	if t.isEnding() {
		return finishingColor
	} else {
		return progressColor
	}
}

func (t *Timer) String(finished bool) string {
	min, sec := t.getMinutesSeconds()
	cf := t.getColorFunc()
	if finished {
		return cf("Finished!")
	}
	res := cf(
		fmt.Sprintf(
			"%s:%s",
			PadLeft(strconv.Itoa(min), "0", 2),
			PadLeft(strconv.Itoa(sec), "0", 2),
		),
	)
	return res
}

func (t *Timer) ProgressElement(state *pb.State, args ...string) string {
	return t.String(state.Current() == state.Total())
}
