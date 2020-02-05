package pomodoro

const (
	stepsCountDefault int = 1
	notifyTimeDefault int = 5 * 60
	workTimeDefault int = 25 * 60
	restTimeDefault int = 5 * 60
	longRestTimeDefault int = 20 * 60
)

type Config struct {
	Steps int
	NotifyTime   int
	WorkTime     int
	RestTime     int
	LongRestTime int
}

func (c *Config) GetStepsCount() int {
	if c.Steps == 0 {
		return stepsCountDefault
	} else {
		return c.Steps
	}
}

func (c *Config) GetNotifyTime() int {
	if c.NotifyTime == -1 {
		return 0
	} else if c.NotifyTime == 0 {
		return notifyTimeDefault
	} else {
		return c.NotifyTime
	}
}

func (c *Config) GetWorkTime() int {
	if c.WorkTime == 0 {
		return workTimeDefault
	} else {
		return c.WorkTime
	}
}

func (c *Config) GetRestTime() int {
	if c.RestTime == 0 {
		return restTimeDefault
	} else {
		return c.RestTime
	}
}

func (c *Config) GetLongRestTime() int {
	if c.LongRestTime == 0 {
		return longRestTimeDefault
	} else {
		return c.LongRestTime
	}
}
