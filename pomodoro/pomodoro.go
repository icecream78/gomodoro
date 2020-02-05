package pomodoro

import (
	"time"
)

var notifyTimeBorder int = 5

type Pomodoro struct {
	stepsCount int
	wTime      int
	timer      Ticker
	stepper    Stepper
	subscribes []Oberver
}

func NewPomodoroTimer(stepsCount int) *Pomodoro {
	return &Pomodoro{
		stepsCount: stepsCount,
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
		return 20 * 60
	} else if currentStep%2 == 0 {
		return 5 * 60
	} else {
		return 25 * 60
	}
}

func (p *Pomodoro) Run() {
	var timer Ticker
	var stepTime int
	var state *State
	stepper := NewStepsCounter(p.stepsCount)

	for !stepper.Finished() {
		stepTime = p.getStepTime(stepper.CurrentStep())
		timer = NewTimer(stepTime, notifyTimeBorder)

		for !timer.Finished() {
			timer.Tick()
			state = &State{
				Step:      stepper.CurrentStep(),
				Progress:  timer.State(),
				TotalStep: p.stepsCount,
				IsEnding:  timer.NeedNotify(),
			}
			go p.Notify(state)
			time.Sleep(1 * time.Second)
		}
		p.Notify(&State{
			Reset: true,
		})
		stepper.NextStep()
		timer.Refresh()
	}
}
