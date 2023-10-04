package enums

import "strings"

const (
	Admin    Enum = "ADMIN"
	Employee Enum = "EMPLOYEE"
)

var mapChargeEnum = map[string]Enum{
	"admin":    Admin,
	"employee": Employee,
}

func StringToChargeEnum(enum string) Enum {
	enum = strings.ToLower(enum)
	return mapChargeEnum[enum]
}
