package clock

import "time"

type Clock interface {
	Now() time.Time
}

type ClockImplementation struct {
}

func (c ClockImplementation) Now() time.Time {
	return time.Now()
}
