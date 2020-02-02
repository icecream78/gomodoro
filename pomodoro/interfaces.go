package pomodoro

import "github.com/cheggaaa/pb/v3"

type Ticker interface {
	Tick()
	Refresh()
	Finished() bool
	WidgetElement
}

type Stepper interface {
	Tick()
	Refresh()
	Finished() bool
	NextStep()
	WidgetElement
}

type WidgetElement interface {
	ProgressElement(state *pb.State, args ...string) string
}
