package pomodoro

type State struct {
	Reset bool

	Step      int
	TotalStep int

	Progress         int
	Finished         bool
	MakeNotification bool
}
