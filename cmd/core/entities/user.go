package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type User struct {
	core.EntitiesBase
	UserName string
	Password string
	Inactive bool
}
