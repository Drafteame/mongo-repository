package clock

import "time"

type Clock struct {
	utc bool
}

func New() Clock {
	return Clock{}
}

func (c Clock) ForceUTC() Clock {
	c.utc = true
	return c
}

func (c Clock) Now() time.Time {
	now := time.Now()

	if c.utc {
		now = now.UTC()
	}

	return now.Truncate(time.Millisecond)
}
