package pomodoro

import (
	"time"

	"github.com/cheggaaa/pb/v3"
)

type Pomodoro struct {
	stepsCount int
	tmpl       string
	wTime      int
	timer      Ticker
	stepper    Stepper
	fns        []Ticker
}

func NewPomodoroTimer(stepsCount int, template string) *Pomodoro {
	return &Pomodoro{
		stepsCount: stepsCount,
		tmpl:       template,
	}
}

func (p *Pomodoro) RegisterTick(fn Ticker) {
	p.fns = append(p.fns, fn)
}

func (p *Pomodoro) runTickers() {
	var fn Ticker
	for i := 0; i < len(p.fns); i++ {
		fn = p.fns[i]
		fn.Tick()
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
	stepper := NewStepsCounter(p.stepsCount)
	var stepTime int

	for !stepper.Finished() {
		stepTime = p.getStepTime(stepper.CurrentStep())
		timer = NewTimer(stepTime, 1*60)

		pb.RegisterElement("timer", timer, true)
		pb.RegisterElement("steps", stepper, true)

		for !timer.Finished() {
			timer.Tick()
			go p.runTickers()
			time.Sleep(1 * time.Second)
		}
		stepper.NextStep()
	}
}
