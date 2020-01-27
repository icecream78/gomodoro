package pomodoro

import "github.com/cheggaaa/pb/v3"

import "time"

type Ticker interface {
	Tick()
	Finished() bool
	Refresh()
}

type Stepper interface {
	NextStep()
	Finished() bool
}

type Widget struct {
	bar     *pb.ProgressBar
	timer   Ticker
	stepper Stepper
	wTime   int
}

func NewWidget(wTime int, timer Ticker, stepper Stepper) *Widget {
	tmpl := `{{ red "Work time:" }} {{bar . "[" "=" "=>" "_" "]"}} {{ wtimer . }} {{ steps . }}`
	bar := pb.ProgressBarTemplate(tmpl).Start64(int64(wTime))

	return &Widget{
		bar:     bar,
		timer:   timer,
		stepper: stepper,
		wTime:   wTime,
	}
}

func (w *Widget) Run() {
	for !w.stepper.Finished() {
		for !w.timer.Finished() {
			w.timer.Tick()
			w.bar.Increment()
			time.Sleep(time.Second)
		}
		w.timer.Refresh()
		w.bar.Finish()
		w.stepper.NextStep()
		tmpl := `{{ red "Work time:" }} {{bar . "[" "=" "=>" "_" "]"}} {{ wtimer . }} {{ steps . }}`
		w.bar = pb.ProgressBarTemplate(tmpl).Start64(int64(w.wTime))
	}
}
