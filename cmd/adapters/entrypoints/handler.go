package entrypoints

import (
	"encoding/json"
	"errors"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	internal_jwt "github.com/fedeveron01/golang-base/cmd/internal/jwt"
	"net/http"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type HandlerBase struct {
	gateways.SessionGateway
}

type MessageResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (h *HandlerBase) WriteResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	response := MessageResponse{
		Message: message,
		Status:  status,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *HandlerBase) WriteUnauthorized(w http.ResponseWriter) {
	h.WriteResponse(w, errors.New("unauthorized").Error(), http.StatusUnauthorized)
}

func (h *HandlerBase) WriteInternalServerError(w http.ResponseWriter, err error) {
	h.WriteResponse(w, err.Error(), http.StatusInternalServerError)
}

func (h *HandlerBase) GetSessionId(r *http.Request) (float64, error) {
	token := r.Header.Get("X-Auth-Token")
	claims, err := internal_jwt.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return claims.SessionId, nil
}

func (h *HandlerBase) IsAuthorized(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("X-Auth-Token")
	if token == "" {
		h.WriteUnauthorized(w)
		return false
	} else {
		claims, err := internal_jwt.ParseToken(token)
		if err != nil {
			h.WriteUnauthorized(w)
			return false
		}
		if claims.Role == "none" {
			h.WriteUnauthorized(w)
			return false
		}

		// validate if session is active

		if claims.SessionId == 0 {
			h.WriteUnauthorized(w)
			return false
		}

		expired := h.SessionGateway.SessionIsExpired(claims.SessionId)
		if expired {
			h.WriteUnauthorized(w)
			return false
		}

		return true
	}
}

func (h *HandlerBase) IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("X-Auth-Token")
	if token == "" {
		h.WriteUnauthorized(w)
		return false
	} else {
		claims, err := internal_jwt.ParseToken(token)
		if err != nil {
			h.WriteUnauthorized(w)
			return false
		}
		if claims.Role != "admin" {
			h.WriteUnauthorized(w)
			return false
		}

		return true
	}
}
