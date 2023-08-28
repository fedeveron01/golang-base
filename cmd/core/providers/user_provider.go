package providers

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
)

type UserProvider interface {
	CreateUser(user entities.User) error
	FindUserByUsernameAndPassword(username string, password string) (entities.User, error)
	FindUserByUsername(username string) entities.User
	UpdateUser(user entities.User) error
	DeleteUser(id string) error
}
