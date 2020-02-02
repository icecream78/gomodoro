package pomodoro

import (
	"time"

	"github.com/cheggaaa/pb/v3"
)

type App struct {
	bar      *pb.ProgressBar
	timer    Ticker
	stepper  Stepper
	wTime    int
	template string
}

func NewWidget(tmpl string, wTime int, timer Ticker, stepper Stepper) *App {
	bar := pb.ProgressBarTemplate(tmpl).Start64(int64(wTime))

	return &App{
		bar:      bar,
		timer:    timer,
		stepper:  stepper,
		wTime:    wTime,
		template: tmpl,
	}
}

func (w *App) Run() {
	for !w.stepper.Finished() {
		for !w.timer.Finished() {
			w.timer.Tick()
			w.bar.Increment()
			time.Sleep(time.Second)
		}
		w.timer.Refresh()
		w.bar.Finish()
		w.stepper.NextStep()
		w.bar = pb.ProgressBarTemplate(w.template).Start64(int64(w.wTime))
	}
}
