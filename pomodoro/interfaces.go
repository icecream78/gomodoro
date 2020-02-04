package pomodoro

type Observable interface {
	Subscribe(sub Oberver)
	Unsubscribe(sub Oberver)
	Notify(state State)
}

type Oberver interface {
	Update(state *State)
}

type Ticker interface {
	Tick()
	Refresh()
	Finished() bool
	NeedNotify() bool
	State() int
}

type Stepper interface {
	Tick()
	Refresh()
	Finished() bool
	NextStep()
}
