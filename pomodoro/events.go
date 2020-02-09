package pomodoro

type Event int

const (
	Init Event = iota
	PreStepHook
	Progress
	PostStepHook
	PostHook
)

func (e Event) String() string {
	switch e {
	case Init:
		return "init"
	case PreStepHook:
		return "pre hook"
	case Progress:
		return "progress"
	case PostHook:
		return "post hook"
	case PostStepHook:
		return "post step hook"
	}
	return ""
}

var (
	AllEvents []Event = []Event{
		Init, PreStepHook, Progress, PostStepHook, PostHook,
	}
)
