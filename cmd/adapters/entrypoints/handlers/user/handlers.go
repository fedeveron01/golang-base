package user_handler

import (
	"encoding/json"
	"errors"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	internal_jwt "github.com/fedeveron01/golang-base/cmd/internal/jwt"
	"github.com/fedeveron01/golang-base/cmd/usecases/user"
	"io"
	"net/http"
)

type UserHandlerInterface interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	entrypoints.HandlerBase
	userUseCase user_usecase.UserUseCase
}

func NewUserHandler(sessionGateway gateways.SessionGateway, userUseCase user_usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		userUseCase: userUseCase,
	}
}

// Handle api/user/signup
func (p *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if !p.IsAdmin(w, r) {
		return
	}

	reqBody, _ := io.ReadAll(r.Body)
	var userRequest CreateUserRequest
	json.Unmarshal(reqBody, &userRequest)
	// call use case and convert request to entities

	err := p.userUseCase.CreateUser(entities.User{
		UserName: userRequest.UserName,
		Password: userRequest.Password,
	}, entities.Employee{
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		DNI:      userRequest.DNI,
		Charge: entities.Charge{
			EntitiesBase: core.EntitiesBase{
				ID: uint(userRequest.ChargeId),
			},
		},
	})
	if err != nil {
		p.WriteInternalServerError(w, err)
		return
	}

	p.WriteResponse(w, "user created", http.StatusCreated)

}

// Handle api/user/login
func (p *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
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
			p.WriteUnauthorized(w)
			return
		}
		p.WriteInternalServerError(w, err)
		return
	}
	claims, _ := internal_jwt.ParseToken(token)

	tokenResponse := TokenResponse{Token: token, EmployeeId: claims.EmployeeId, Charge: claims.Role}
	json.NewEncoder(w).Encode(tokenResponse)
}

// Handle api/user/logout
func (p *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	sessionId, err := p.GetSessionId(r)
	if err != nil {
		p.WriteUnauthorized(w)
		return
	}

	err = p.userUseCase.LogoutUser(sessionId)

	if err != nil {
		p.WriteInternalServerError(w, err)
		return
	}

}
