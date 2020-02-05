package widget

import (
	"github.com/fatih/color"
)

var (
	progressColor  = color.New(color.FgRed, color.Bold).SprintFunc()
	finishingColor = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	doneColor      = color.New(color.FgGreen, color.Bold).SprintFunc()
)
