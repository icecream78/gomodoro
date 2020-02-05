package pomodoro

import (
	"time"
)

type Pomodoro struct {
	*Config
	timer      Ticker
	stepper    Stepper
	subscribes []Oberver
}

func NewPomodoroTimer(c *Config) *Pomodoro {
	return &Pomodoro{
		Config: c,
	}
}

func (p *Pomodoro) Subscribe(sub Oberver) {
	p.subscribes = append(p.subscribes, sub)
}

// TODO: add method for removing subscriber from list
func (p *Pomodoro) Unsubscribe(sub Oberver) {
}

func (p *Pomodoro) Notify(state *State) {
	var fn Oberver
	for i := 0; i < len(p.subscribes); i++ {
		fn = p.subscribes[i]
		fn.Update(state)
	}
}

func (p *Pomodoro) getStepTime(currentStep int) int {
	if currentStep%5 == 0 {
		return p.GetLongRestTime()
	} else if currentStep%2 == 0 {
		return p.GetRestTime()
	} else {
		return p.GetWorkTime()
	}
}

func (p *Pomodoro) Run() {
	var timer Ticker
	var stepTime int
	var state *State
	stepper := NewStepsCounter(p.GetStepsCount())

	for !stepper.Finished() {
		stepTime = p.getStepTime(stepper.CurrentStep())
		timer = NewTimer(stepTime, p.GetNotifyTime())
		p.Notify(&State{
			Reset:     true,
			TotalTime: stepTime,
		})

		for !timer.Finished() {
			timer.Tick()
			state = &State{
				Step:      stepper.CurrentStep(),
				TotalStep: p.GetStepsCount(),
				Progress:  timer.State(),
				IsEnding:  timer.NeedNotify(),
			}
			go p.Notify(state)
			time.Sleep(1 * time.Second)
		}
		stepper.NextStep()
		timer.Refresh()
	}
}
