package enums

import "strings"

type Enum string // Enum is a string type

func (e Enum) String() string {
	return strings.ToLower(string(e))
}
