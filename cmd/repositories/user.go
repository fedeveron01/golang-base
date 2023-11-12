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

func (r *UserRepository) CreateCompleteUserWithEmployee(
	user gateway_entities.User,
	chargeID uint,
	employee gateway_entities.Employee) (gateway_entities.User, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user)
		employee.UserId = user.ID
		employee.ChargeId = chargeID

		res := tx.Create(&employee)
		if res.Error != nil {
			return res.Error
		}
		return nil

	})
	if err != nil {
		return gateway_entities.User{}, err
	}
	user = r.FindUserByUsername(user.UserName)
	return user, nil

}

func (r *UserRepository) CreateUser(user gateway_entities.User) (gateway_entities.User, error) {
	var userDB gateway_entities.User
	id := r.db.Create(&user)
	if id.RowsAffected == 0 {
		return gateway_entities.User{}, errors.New("user not created")
	}
	r.db.Where("user_name = ?", user.UserName).First(&userDB)

	return userDB, nil
}

func (r *UserRepository) FindUserById(id int64) (user gateway_entities.User, err error) {
	res := r.db.Where("id = ?", id).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (r *UserRepository) FindUserByUsername(username string) gateway_entities.User {
	var user gateway_entities.User
	res := r.db.Where("user_name = ?", username).First(&user)
	if res.Error != nil {
		return gateway_entities.User{}
	}
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

func (r *UserRepository) UpdateUser(user gateway_entities.User) (gateway_entities.User, error) {
	res := r.db.Save(&user)
	if res.Error != nil {
		return gateway_entities.User{}, res.Error
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id string) error {
	r.db.Delete(&gateway_entities.User{}, id)
	return nil
}
