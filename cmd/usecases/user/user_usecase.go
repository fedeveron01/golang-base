package user_usecase

import (
	"errors"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	internal_jwt "github.com/fedeveron01/golang-base/cmd/internal/jwt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUseCase interface {
	CreateUser(user entities.User, employee entities.Employee) (string, error)
	UpdateUser(user entities.User) error
	DeleteUser(id string) error
	LoginUser(username string, password string) (string, error)
}

type UserGateway interface {
	CreateUser(user entities.User) error
	FindUserByUsernameAndPassword(username string, password string) (entities.User, error)
	FindUserByUsername(username string) entities.User
	UpdateUser(user entities.User) error
	DeleteUser(id string) error
}

type SessionGateway interface {
	CreateSession(session entities.Session) error
	FindAll() ([]entities.Session, error)
	UpdateSession(session entities.Session) error
	DeleteSession(id string) error
}

type EmployeeGateway interface {
	CreateEmployee(employee entities.Employee) error
	FindEmployeeByUserId(id uint) (entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	UpdateEmployee(employee entities.Employee) error
	DeleteEmployee(id string) error
}

type Implementation struct {
	userGateway     UserGateway
	sessionGateway  SessionGateway
	employeeGateway EmployeeGateway
}

func NewUserUseCase(userGateway UserGateway,
	sessionGateway SessionGateway,
	employeeGateway EmployeeGateway) *Implementation {
	return &Implementation{
		userGateway:     userGateway,
		sessionGateway:  sessionGateway,
		employeeGateway: employeeGateway,
	}
}

func (i *Implementation) CreateUser(user entities.User, employee entities.Employee) (string, error) {
	if user.UserName == "" || user.Password == "" {
		return "", errors.New("username or password is empty")
	}
	user.Password = encryptPassword(user.Password)

	userRepeated := i.userGateway.FindUserByUsername(user.UserName)
	if userRepeated.ID != 0 {
		return "", errors.New("username already exists")
	}
	err := i.userGateway.CreateUser(user)
	if err != nil {
		return "", err
	}
	employee.User = user
	err = i.employeeGateway.CreateEmployee(employee)
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
	user := i.userGateway.FindUserByUsername(username)
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
	err := i.sessionGateway.CreateSession(session)
	if err != nil {
		return "", err
	}
	// get employee
	employee, err := i.employeeGateway.FindEmployeeByUserId(user.ID)
	if err != nil {
		return "", err
	}
	return generateToken(employee)
}
func (i *Implementation) UpdateUser(user entities.User) error {
	return i.userGateway.UpdateUser(user)
}
func (i *Implementation) DeleteUser(id string) error {
	return i.userGateway.DeleteUser(id)
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
		EmployeeId: float64(employee.ID),
		Role:       role,
	}

	return t.SignedString([]byte("test"))
}
