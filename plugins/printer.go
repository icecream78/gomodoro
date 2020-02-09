package plugins

import (
	"fmt"

	"github.com/icecream78/gomodoro/pomodoro"
)

type Printer struct {
	printText string
}

func NewPrinter(text string) *Printer {
	return &Printer{printText: text}
}

func (p *Printer) Update(state *pomodoro.State) {
	fmt.Printf("Text: %s by hook %s\n", p.printText, state.Event.String())
}
