package calculate_age

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type CalculateAgeMock struct {
	mock.Mock
}

func (c *CalculateAgeMock) Now() time.Time {
	args := c.Called()
	return args.Get(0).(time.Time)
}
