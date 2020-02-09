package pomodoro

type State struct {
	Event Event

	Step      int
	TotalStep int

	Progress  int
	TotalTime int
	Finished  bool
	IsEnding  bool
}
