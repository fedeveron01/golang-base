package enums

import "strings"

const (
	Input  Enum = "INPUT"
	Output Enum = "OUTPUT"
)

var mapMovementTypeEnum = map[string]Enum{
	"input":  Input,
	"output": Output,
}

func StringToMovementTypeEnum(enum string) Enum {
	enum = strings.ToLower(enum)
	return mapMovementTypeEnum[enum]
}
