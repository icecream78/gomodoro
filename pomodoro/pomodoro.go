package pomodoro

import (
	"time"
)

type Pomodoro struct {
	*Config
	timer      Ticker
	stepper    Stepper
	subscribes map[Event][]Oberver
}

func NewPomodoroTimer(c *Config) *Pomodoro {
	return &Pomodoro{
		Config:     c,
		subscribes: make(map[Event][]Oberver),
	}
}

func (p *Pomodoro) Subscribe(sub Oberver) {
	p.SubscribeEvent(sub, AllEvents...)
}

func (p *Pomodoro) SubscribeEvent(sub Oberver, events ...Event) {
	for _, event := range events {
		p.subscribes[event] = append(p.subscribes[event], sub)
	}
}

func (p *Pomodoro) Unsubscribe(sub Oberver) {
	p.UnsubscribeEvent(sub, Progress)
}

// TODO: add method for removing subscriber from list
func (p *Pomodoro) UnsubscribeEvent(sub Oberver, event Event) {
}

func (p *Pomodoro) Notify(state *State) {
	p.notifyEvent(state.Event, state)
}

func (p *Pomodoro) notifyEvent(event Event, state *State) {
	var fn Oberver
	coll := p.subscribes[event]
	for i := 0; i < len(coll); i++ {
		fn = coll[i]
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
		timer = NewTimer(stepTime, p.GetNotifyPercent())
		p.Notify(&State{
			Event:     PreStepHook,
			TotalTime: stepTime,
		})

		for !timer.Finished() {
			timer.Tick()
			state = &State{
				Event: Progress,

				Step:      stepper.CurrentStep(),
				TotalStep: p.GetStepsCount(),
				Progress:  timer.State(),
				IsEnding:  timer.NeedNotify(),
			}
			go p.Notify(state)
			time.Sleep(1 * time.Second)
		}
		p.Notify(&State{
			Event:     PostStepHook,
			IsEnding:  true,
			Step:      stepper.CurrentStep(),
			TotalStep: p.GetStepsCount(),
			Progress:  timer.State(),
		})

		stepper.NextStep()
		timer.Refresh()
	}
	p.Notify(&State{
		Event: PostHook,
	})
}
