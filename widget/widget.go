package widget

import (
	"fmt"
	"strconv"

	"github.com/icecream78/gomodoro/pomodoro"

	"github.com/cheggaaa/pb/v3"
)

const (
	ZeroedTime      int = 0
	ZeroedStepper   int = 1
	DefaultShowTime int = 100
)

type colorFunc func(a ...interface{}) string

// NewBar returns new copy of widget that is shown in terminal
func NewBar(template string) *Widget {
	w := Widget{
		template: template,
		showTime: DefaultShowTime,
	}
	return &w
}

type Widget struct {
	template string
	showTime int
	bar      *pb.ProgressBar
}

func (w *Widget) Tick() {
	w.bar.Increment()
}

func (w *Widget) InitBar(barTime int) {
	w.bar = pb.ProgressBarTemplate(w.template).Start(barTime)
	return
}

func (w *Widget) Update(state *pomodoro.State) {
	if state.Event == pomodoro.PreStepHook {
		if w.bar != nil {
			w.bar.Finish()
		}
		w.InitBar(state.TotalTime)
		return
	}
	w.bar.Increment()

	ts := w.renderTimer(state.Progress, state.IsEnding)
	ss := w.formatSteps(state.Step, state.TotalStep)

	w.bar.Set("timer", ts)
	w.bar.Set("steps", ss)
	w.bar.Set("newline", "1")

	// TODO: solve trouble with non printing char in template
	if state.Event == pomodoro.PostStepHook {
		w.bar.Set("newline", "\n")
	}
}

func (w *Widget) getTimerColorFunc(now int, isEnding bool) colorFunc {
	if now == 0 {
		return doneColor
	} else if isEnding {
		return finishingColor
	} else {
		return progressColor
	}
}

func (w *Widget) renderTimer(time int, isEnding bool) string {
	cf := w.getTimerColorFunc(time, isEnding)
	res := cf(w.formatTime(time))
	return res
}

func (w *Widget) formatTime(time int) string {
	min, sec := getMinutesSeconds(time)
	if min == 0 && sec == 0 {
		return "Finished!"
	}
	return fmt.Sprintf(
		"%s:%s",
		padLeft(strconv.Itoa(min), "0", 2),
		padLeft(strconv.Itoa(sec), "0", 2),
	)
}

func (w *Widget) formatSteps(current int, total int) string {
	if total == 1 {
		return ""
	}
	return fmt.Sprintf("%d/%d", current, total)
}
