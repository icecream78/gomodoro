package plugins

import (
	"github.com/gen2brain/beeep"
	"github.com/icecream78/gomodoro/pomodoro"
)

type Notification struct {
	title     string
	printText string
}

func NewNotification(title string, text string) *Notification {
	return &Notification{printText: text, title: title}
}

func (n *Notification) Update(state *pomodoro.State) {
	if err := beeep.Notify(n.title, n.printText, "assets/information.png"); err != nil {
		panic(err)
	}
}
