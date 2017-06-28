package clock

import (
	"time"
)

const testVersion = 4

type Clock string

var dt time.Time

func New(hour, minute int) Clock {
	dt = time.Date(1, 1, 1, hour, minute, 0, 0, time.UTC)
	return Clock(dt.Format("15:04"))
}

func (c Clock) String() string {
	return string(c)
}

func (c Clock) Add(minutes int) Clock {
	return Clock(dt.Add(time.Minute * time.Duration(minutes)).Format("15:04"))
}
