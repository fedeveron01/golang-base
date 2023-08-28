package providers

import "github.com/fedeveron01/golang-base/cmd/core/entities"

type SessionProvider interface {
	CreateSession(session entities.Session) error
	FindAll() ([]entities.Session, error)
	UpdateSession(session entities.Session) error
	DeleteSession(id string) error
}
