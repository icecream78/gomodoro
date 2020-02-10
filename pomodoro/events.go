package pomodoro

type Event int

const (
	Init Event = iota
	PreStepHook
	Progress
	PostStepHook
	PostHook
	WorkStepStart
	RestStepStart
	LongRestStepStart
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
	case WorkStepStart:
		return "work step"
	case RestStepStart:
		return "rest step"
	case LongRestStepStart:
		return "long rest step"
	}
	return ""
}

var (
	AllEvents []Event = []Event{
		Init, PreStepHook, Progress, PostStepHook, PostHook,
	}
)
