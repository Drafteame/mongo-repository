package clock

import "time"

type TestClock struct {
	now time.Time
}

func NewTest(now time.Time) TestClock {
	return TestClock{now: now}
}

func (c TestClock) ForceUTC() TestClock {
	c.now = c.now.UTC()
	return c
}

func (c TestClock) Now() time.Time {
	return c.now.Truncate(time.Millisecond)
}
