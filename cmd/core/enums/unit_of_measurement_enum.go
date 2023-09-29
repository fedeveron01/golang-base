package enums

import "strings"

const (
	Liters    Enum = "LITERS"
	Meters    Enum = "METERS"
	Units     Enum = "UNITS"
	Kilograms Enum = "KILOGRAMS"
)

var mapUnitOfMeasurementEnum = map[string]Enum{
	"liters":    Liters,
	"meters":    Meters,
	"units":     Units,
	"kilograms": Kilograms,
}

func StringToUnitOfMeasurementEnum(enum string) Enum {
	enum = strings.ToLower(enum)
	return mapUnitOfMeasurementEnum[enum]
}
