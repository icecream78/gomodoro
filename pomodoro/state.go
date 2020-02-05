package pomodoro

type State struct {
	Reset bool

	Step      int
	TotalStep int

	Progress  int
	TotalTime int
	Finished  bool
	IsEnding  bool
}
