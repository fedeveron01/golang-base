package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type Session struct {
	core.EntitiesBase
	User User
}
