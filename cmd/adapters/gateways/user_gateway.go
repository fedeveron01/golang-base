package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type UserGatewayImpl struct {
	userRepository repositories.UserRepository
}

func NewUserGateway(userRepository repositories.UserRepository) *UserGatewayImpl {
	return &UserGatewayImpl{
		userRepository: userRepository,
	}
}

func (i *UserGatewayImpl) CreateUser(user entities.User) error {
	userDB := gateway_entities.User{
		UserName: user.UserName,
		Password: user.Password,
	}
	err := i.userRepository.CreateUser(userDB)
	if err != nil {
		return err
	}
	return nil
}

func (i *UserGatewayImpl) FindUserByUsernameAndPassword(username string, password string) (entities.User, error) {
	userDB, err := i.userRepository.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		return entities.User{}, err
	}
	user := entities.User{
		EntitiesBase: core.EntitiesBase{
			ID: userDB.ID,
		},
		UserName: userDB.UserName,
		Password: userDB.Password,
	}
	return user, nil
}

func (i *UserGatewayImpl) FindUserByUsername(username string) entities.User {
	userDB := i.userRepository.FindUserByUsername(username)
	user := entities.User{
		EntitiesBase: core.EntitiesBase{
			ID: userDB.ID,
		},
		UserName: userDB.UserName,
		Password: userDB.Password,
	}
	return user
}

func (i *UserGatewayImpl) UpdateUser(user entities.User) error {
	userDB := gateway_entities.User{
		UserName: user.UserName,
		Password: user.Password,
	}
	return i.userRepository.UpdateUser(userDB)
}

func (i *UserGatewayImpl) DeleteUser(id string) error {
	return i.userRepository.DeleteUser(id)
}
