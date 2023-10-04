package enums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToMovementTypeEnumWhenStringIsNotValidShouldReturnEmptyEnum(t *testing.T) {
	result := StringToMovementTypeEnum("invalid")
	assert.Equal(t, result, Enum(""))
}

func TestStringToMovementTypeEnumWhenStringIsValidShouldReturnEnum(t *testing.T) {
	result := StringToMovementTypeEnum("input")
	assert.Equal(t, result, Input)
}

func TestStringToMovementTypeEnumWhenStringIsValidButNotLowercaseShouldReturnEnum(t *testing.T) {
	result := StringToMovementTypeEnum("INPUT")
	assert.Equal(t, result, Input)
}

func TestStringToMovementTypeEnumWhenStringIsValidButHasUppercaseAndLowercaseShouldReturnEnum(t *testing.T) {
	result := StringToMovementTypeEnum("InPuT")
	assert.Equal(t, result, Input)
}
