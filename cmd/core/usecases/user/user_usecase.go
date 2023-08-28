package user_usecase

import (
	"errors"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/providers"
	internal_jwt "github.com/fedeveron01/golang-base/cmd/internal/jwt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUsecase interface {
	CreateUser(user entities.User) error
	FindUserByUsernameAndPassword(username string, password string) (entities.User, error)
	UpdateUser(user entities.Material) error
	DeleteUser(id string) error
}
type Implementation struct {
	userProvider     providers.UserProvider
	sessionProvider  providers.SessionProvider
	employeeProvider providers.EmployeeProvider
}

func NewUserUsecase(userProvider providers.UserProvider,
	sessionProvider providers.SessionProvider,
	employeeProvider providers.EmployeeProvider) *Implementation {
	return &Implementation{
		userProvider:     userProvider,
		sessionProvider:  sessionProvider,
		employeeProvider: employeeProvider,
	}
}

func (i *Implementation) CreateUser(user entities.User) (string, error) {
	if user.UserName == "" || user.Password == "" {
		return "", errors.New("username or password is empty")
	}
	user.Password = encryptPassword(user.Password)

	userRepeated := i.userProvider.FindUserByUsername(user.UserName)
	if userRepeated.ID != 0 {
		return "", errors.New("username already exists")
	}
	err := i.userProvider.CreateUser(user)
	if err != nil {
		return "", err
	}
	return i.LoginUser(user.UserName, user.Password)
}

func encryptPassword(password string) string {
	encriptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encriptedPassword)
}

func isCorrectPassword(encryptedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (i *Implementation) LoginUser(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("username or password is empty")
	}
	user := i.userProvider.FindUserByUsername(username)
	if user.ID == 0 {
		return "", errors.New("user not found")
	}
	if !isCorrectPassword(user.Password, password) {
		return "", errors.New("username or password is incorrect")
	}
	// create session
	session := entities.Session{
		User: user,
	}
	err := i.sessionProvider.CreateSession(session)
	if err != nil {
		return "", err
	}
	// get employee
	employee, err := i.employeeProvider.FindEmployeeByUserId(user.ID)
	if err != nil {
		return "", err
	}
	return generateToken(employee)
}
func (i *Implementation) UpdateUser(user entities.User) error {
	return i.userProvider.UpdateUser(user)
}
func (i *Implementation) DeleteUser(id string) error {
	return i.userProvider.DeleteUser(id)
}

func generateToken(employee entities.Employee) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	var role string
	if employee.ID == 0 {
		role = "none"
	} else {
		role = employee.Charge.Name
	}
	t.Claims = &internal_jwt.Claims{
		StandardClaims: &jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		TokenType:  "level1",
		EmployeeId: employee.ID,
		Role:       role,
	}

	return t.SignedString([]byte("test"))
}
