package user_handler

import (
	"encoding/json"
	"errors"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	internal_jwt "github.com/fedeveron01/golang-base/cmd/internal/jwt"
	"github.com/fedeveron01/golang-base/cmd/usecases/user"
	"io"
	"net/http"
)

type CreateUserHandler struct {
	entrypoints.HandlerBase
	userUseCase user_usecase.UserUseCase
}

func NewCreateUserHandler(userUseCase user_usecase.UserUseCase) CreateUserHandler {
	return CreateUserHandler{
		HandlerBase: entrypoints.HandlerBase{},
		userUseCase: userUseCase,
	}
}

// Handle api/user/signup
func (p CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var userRequest CreateUserRequest
	json.Unmarshal(reqBody, &userRequest)
	// call use case and convert request to entities
	token, err := p.userUseCase.CreateUser(entities.User{
		UserName: userRequest.UserName,
		Password: userRequest.Password,
	}, entities.Employee{
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		DNI:      userRequest.DNI,
		Charge: entities.Charge{
			Name: userRequest.Charge,
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	tokenResponse := TokenResponse{Token: token}
	json.NewEncoder(w).Encode(tokenResponse)

}

type LoginUserHandler struct {
	entrypoints.HandlerBase
	userUseCase user_usecase.UserUseCase
}

func NewLoginUserHandler(userUseCase user_usecase.UserUseCase) LoginUserHandler {
	return LoginUserHandler{
		HandlerBase: entrypoints.HandlerBase{},
		userUseCase: userUseCase,
	}
}

// Handle api/user/login
func (p LoginUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var loginRequest LoginRequest

	err := json.Unmarshal(reqBody, &loginRequest)

	if err != nil {
		p.WriteInternalServerError(w, err)
		return
	}
	token, err := p.userUseCase.LoginUser(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		if errors.Is(err, core_errors.ErrInactiveUser) {
			p.WriteUnauthorizedError(w, err)
			return
		}
		p.WriteInternalServerError(w, err)
		return
	}
	claims, _ := internal_jwt.ParseToken(token)

	tokenResponse := TokenResponse{Token: token, EmployeeId: claims.EmployeeId, Charge: claims.Role}
	json.NewEncoder(w).Encode(tokenResponse)
}
