package calculate_age

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/internal/clock"
)

type CalculateAgeUseCase interface {
	CalculateAge(person entities.Person) int
}

type Implementation struct {
}

func (c Implementation) CalculateAge(person entities.Person) int {
	clock := clock.ClockImplementation{}
	return person.CalculateAge(clock)
}
