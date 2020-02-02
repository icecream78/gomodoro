package pomodoro

import (
	"github.com/cheggaaa/pb/v3"
)

func NewBar(template string, showTime int) *Widget {
	bar := pb.ProgressBarTemplate(template).Start(showTime)
	return &Widget{
		template: template,
		showTime: showTime,
		bar:      bar,
	}
}

type Widget struct {
	template string
	showTime int
	bar      *pb.ProgressBar
}

func (w *Widget) RegisterElement(alias string, element WidgetElement) {
	pb.RegisterElement(alias, element, true)
}

func (b *Widget) Tick() {
	b.bar.Increment()
}

// TODO: write proper logic
func (b *Widget) Finished() bool {
	return false
}

// TODO: write proper logic
func (b *Widget) Refresh() {
	return
}

// TODO: remove handler
func (b *Widget) ProgressElement(state *pb.State, args ...string) string {
	return ""
}

func (b *Widget) Run() {
}
