package user_handler

import (
	"encoding/json"
	"fmt"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/usecases/user"
	"io"
	"net/http"
)

type CreateUserHandler struct {
	userUsecase user_usecase.Implementation
}

func NewCreateUserHandler(userUsecase user_usecase.Implementation) CreateUserHandler {
	return CreateUserHandler{
		userUsecase: userUsecase,
	}
}

func (p CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var user entities.User
	json.Unmarshal(reqBody, &user)
	token, err := p.userUsecase.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	tokenResponse := TokenResponse{Token: token}
	json.NewEncoder(w).Encode(tokenResponse)

}

type LoginUserHandler struct {
	userUsecase user_usecase.Implementation
}

func NewLoginUserHandler(userUsecase user_usecase.Implementation) LoginUserHandler {
	return LoginUserHandler{
		userUsecase: userUsecase,
	}
}

func (p LoginUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var loginRequest LoginRequest

	err := json.Unmarshal(reqBody, &loginRequest)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	token, err := p.userUsecase.LoginUser(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	tokenResponse := TokenResponse{Token: token}
	json.NewEncoder(w).Encode(tokenResponse)
}
