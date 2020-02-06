package pomodoro

const (
	minutesInSeconds        = 60
	workTimeDefault     int = 25 * minutesInSeconds
	restTimeDefault     int = 5 * minutesInSeconds
	longRestTimeDefault int = 20 * minutesInSeconds

	stepsCountDefault    int = 1
	notifyPercentDefault int = 25
)

type Config struct {
	Steps        int
	Notify       int
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

func (c *Config) GetNotifyPercent() int {
	if c.Notify == -1 {
		return 0
	} else if c.Notify == 0 {
		return notifyPercentDefault
	} else {
		return c.Notify
	}
}

func (c *Config) GetWorkTime() int {
	if c.WorkTime == 0 {
		return workTimeDefault
	} else {
		return c.WorkTime * minutesInSeconds
	}
}

func (c *Config) GetRestTime() int {
	if c.RestTime == 0 {
		return restTimeDefault
	} else {
		return c.RestTime * minutesInSeconds
	}
}

func (c *Config) GetLongRestTime() int {
	if c.LongRestTime == 0 {
		return longRestTimeDefault
	} else {
		return c.LongRestTime * minutesInSeconds
	}
}
