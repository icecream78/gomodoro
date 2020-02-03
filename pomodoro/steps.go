package pomodoro

import (
	"sync"
)

func NewStepsCounter(steps int) *StepsCounter {
	return &StepsCounter{
		steps:   steps,
		current: 1,
	}
}

type StepsCounter struct {
	steps   int
	current int
	mx      sync.RWMutex
}

func (s *StepsCounter) SetStepsCount(i int) {
	s.mx.Lock()
	s.steps = i
	s.mx.Unlock()
	return
}

func (s *StepsCounter) NextStep() {
	s.mx.Lock()
	s.current += 1
	s.mx.Unlock()
	return
}

func (s *StepsCounter) CurrentStep() int {
	s.mx.Lock()
	curr := s.current
	s.mx.Unlock()
	return curr
}

func (s *StepsCounter) Finished() bool {
	return s.CurrentStep() == s.steps+1
}
