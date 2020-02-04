package widget

import (
	"fmt"
	"strconv"

	"github.com/icecream78/gomodoro/pomodoro"

	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
)

type colorFunc func(a ...interface{}) string

var (
	progressColor  = color.New(color.FgRed, color.Bold).SprintFunc()
	finishingColor = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	finishedColor  = color.New(color.FgGreen, color.Bold).SprintFunc()
)

func PadLeft(original string, padSymb string, length int) string {
	if length <= 0 || len(original) >= length {
		return original
	}
	s := original
	for i := 0; i < length-len(original); i++ {
		s = padSymb + s
	}
	return s
}

func NewBar(template string, showTime int) *Widget {
	bar := pb.ProgressBarTemplate(template).Start(showTime)
	bar.Set("timer", "00:00")
	bar.Set("steps", "")
	return &Widget{
		template:     template,
		showTime:     showTime,
		bar:          bar,
		notifyBorder: 5, // TODO: remove hardcode
	}
}

type Widget struct {
	template     string
	showTime     int
	bar          *pb.ProgressBar
	notifyBorder int
}

func (b *Widget) Tick() {
	b.bar.Increment()
}

// TODO: write proper logic
func (b *Widget) Refresh() {
	bar := pb.ProgressBarTemplate(b.template).Start(b.showTime)
	bar.Set("timer", "00:00")
	bar.Set("steps", "")
	b.bar = bar
	return
}

func (b *Widget) Update(state *pomodoro.State) {
	b.bar.Increment()
	if state.Reset {
		b.bar.Finish()
		b.Refresh()
		return
	}

	ts := b.RenderTimer(state.Progress)
	ss := b.RenderSteps(state.Step, state.TotalStep)

	b.bar.Set("timer", ts)
	b.bar.Set("steps", ss)
}

func (t *Widget) getMinutesSeconds(now int) (min int, sec int) {
	if now > 0 {
		sec = now % 60
		min = (now - sec) / 60
	}
	return
}

func (t *Widget) isEnding(now int) bool {
	return now <= t.notifyBorder
}

func (t *Widget) getColorFunc(now int) colorFunc {
	if t.isEnding(now) {
		return finishingColor
	} else {
		return progressColor
	}
}

func (t *Widget) RenderTimer(now int) string {
	min, sec := t.getMinutesSeconds(now)
	cf := t.getColorFunc(now)
	if min == 0 && sec == 0 {
		return cf("Finished!")
	}
	res := cf(
		fmt.Sprintf(
			"%s:%s",
			PadLeft(strconv.Itoa(min), "0", 2),
			PadLeft(strconv.Itoa(sec), "0", 2),
		),
	)
	return res
}

func (s *Widget) RenderSteps(current int, total int) string {
	if total == 1 {
		return ""
	}
	return fmt.Sprintf("%d/%d", current, total)
}
