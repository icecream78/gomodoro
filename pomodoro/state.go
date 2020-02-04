package pomodoro

type State struct {
	Step      int
	TotalStep int

	Progress         int
	Finished         bool
	MakeNotification bool
}
