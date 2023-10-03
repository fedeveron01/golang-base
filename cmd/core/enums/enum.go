package enums

import "strings"

type Enum string // Enum is a string type

func (e Enum) String(language string) string {
	if language == "es" {
		return EnumToUnitOfMeasurementStringInSpanish(e)
	}

	return strings.ToLower(string(e))
}
