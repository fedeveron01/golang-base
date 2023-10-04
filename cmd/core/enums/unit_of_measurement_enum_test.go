package enums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToUnitOfMeasurementEnumWhenEnumIsValidShouldReturnEmptyEnum(t *testing.T) {
	result := StringToUnitOfMeasurementEnum("litrs")
	assert.Equal(t, result, Enum(""))
}

func TestStringToUnitOfMeasurementEnumWhenEnumIsValidShouldReturnEnum(t *testing.T) {
	result := StringToUnitOfMeasurementEnum("liters")
	assert.Equal(t, result, Liters)
}

func TestStringToUnitOfMeasurementEnumWhenEnumIsValidButNotLowercaseShouldReturnEnum(t *testing.T) {
	result := StringToUnitOfMeasurementEnum("LITERS")
	assert.Equal(t, result, Liters)
}

func TestStringToUnitOfMeasurementEnumWhenEnumIsValidButHasUppercaseAndLowercaseShouldReturnEnum(t *testing.T) {
	result := StringToUnitOfMeasurementEnum("LiTros")
	assert.Equal(t, result, Liters)
}
