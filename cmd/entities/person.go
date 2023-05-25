package entities

import (
	"time"

	"github.com/fedeveron01/golang-base/cmd/internal/clock"
)

type Person struct {
	Id       int
	Name     string
	LastName string
	BornDate time.Time
}

func (p *Person) CalculateAge(clock clock.Clock) int {
	return clock.Now().Year() - p.BornDate.Year()
}
