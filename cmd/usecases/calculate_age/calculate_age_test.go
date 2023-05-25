package calculate_age

import (
	"testing"
	"time"

	"github.com/fedeveron01/golang-base/cmd/entities"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAgeWhenBornDateIsValidShouldReturnAValidAge(t *testing.T) {
	mockClock := new(CalculateAgeMock)

	//represent today
	mockClock.On("Now").Return(time.Date(2023, 2, 19, 14, 0, 0, 0, time.Local))

	person := entities.Person{BornDate: time.Date(2000, 2, 19, 14, 0, 0, 0, time.Local)}

	expectedAge := 23
	actualAge := person.CalculateAge(mockClock)

	assert.Equal(t, expectedAge, actualAge)
}
