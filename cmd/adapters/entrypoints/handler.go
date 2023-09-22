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

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ToErrorResponse(w http.ResponseWriter, err error, status int) {
	errorResponse := ErrorResponse{
		Message: err.Error(),
		Status:  status,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}

func (h *HandlerBase) WriteUnauthorized(w http.ResponseWriter) {
	ToErrorResponse(w, errors.New("unauthorized"), http.StatusUnauthorized)
}

func (h *HandlerBase) WriteInternalServerError(w http.ResponseWriter, err error) {
	ToErrorResponse(w, err, http.StatusInternalServerError)

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
		h.writeUnauthorized(w)
		return false
	} else {
		claims, err := internal_jwt.ParseToken(token)
		if err != nil {
			h.writeUnauthorized(w)
			return false
		}
		if claims.Role != "admin" {
			h.writeUnauthorized(w)
			return false
		}

		return true
	}
}
