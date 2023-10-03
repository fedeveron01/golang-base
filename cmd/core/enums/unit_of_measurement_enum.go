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

var mapUnitOfMeasurementEnumInSpanish = map[string]Enum{
	"litros":     Liters,
	"metros":     Meters,
	"unidades":   Units,
	"kilogramos": Kilograms,
}

func GetSymbolByUnitOfMeasurementEnum(enum Enum) string {
	switch enum {
	case Liters:
		return "L"
	case Meters:
		return "m"
	case Units:
		return "u"
	case Kilograms:
		return "kg"
	}
	return ""
}
func StringToUnitOfMeasurementEnum(enum string) Enum {
	enum = strings.ToLower(enum)

	res := mapUnitOfMeasurementEnum[enum]
	if res == "" {
		return mapUnitOfMeasurementEnumInSpanish[enum]
	}
	return res
}

func EnumToUnitOfMeasurementStringInSpanish(enum Enum) string {
	switch enum {
	case Liters:
		return "litros"
	case Meters:
		return "metros"
	case Units:
		return "unidades"
	case Kilograms:
		return "kilogramos"
	}
	return ""
}

func GetAllUnitOfMeasurementEnum(language string) []string {
	return []string{Liters.String(language), Meters.String(language), Units.String(language), Kilograms.String(language)}

}
