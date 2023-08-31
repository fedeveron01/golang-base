package repositories

import (
	"errors"
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

func (r UserRepository) CreateUser(user gateway_entities.User) error {
	id := r.db.Create(&user)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *UserRepository) FindUserByUsername(username string) gateway_entities.User {
	var user gateway_entities.User
	r.db.Where("user_name = ?", username).First(&user)
	return user
}

func (r *UserRepository) FindUserByUsernameAndPassword(username string, password string) (gateway_entities.User, error) {
	var user gateway_entities.User
	user.UserName = username
	user.Password = password
	res := r.db.Where("user_name = ? AND password = ?", username, password).First(&user)
	if res.Error != nil {
		return gateway_entities.User{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user gateway_entities.User) error {
	r.db.Save(&user)
	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	r.db.Delete(&gateway_entities.User{}, id)
	return nil
}
