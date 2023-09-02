package user_handler

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/usecases/user"
	"io"
	"net/http"
)

type CreateUserHandler struct {
	userUsercase user_usecase.UserUsecase
}

func NewCreateUserHandler(userUsecase user_usecase.UserUsecase) CreateUserHandler {
	return CreateUserHandler{
		userUsercase: userUsecase,
	}
}

func (p CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var user entities.User
	json.Unmarshal(reqBody, &user)
	token, err := p.userUsercase.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	tokenResponse := TokenResponse{Token: token}
	json.NewEncoder(w).Encode(tokenResponse)

}

type LoginUserHandler struct {
	userUsecase user_usecase.UserUsecase
}

func NewLoginUserHandler(userUsecase user_usecase.UserUsecase) LoginUserHandler {
	return LoginUserHandler{
		userUsecase: userUsecase,
	}
}

func (p LoginUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var loginRequest LoginRequest

	err := json.Unmarshal(reqBody, &loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	token, err := p.userUsecase.LoginUser(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	tokenResponse := TokenResponse{Token: token}
	json.NewEncoder(w).Encode(tokenResponse)
}
