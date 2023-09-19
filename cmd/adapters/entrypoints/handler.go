package entrypoints

import (
	"encoding/json"
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

func (h *HandlerBase) writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Unauthorized")
}

func (h *HandlerBase) WriteInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (h *HandlerBase) WriteUnauthorizedError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(err.Error())
}

func (h *HandlerBase) IsAuthorized(w http.ResponseWriter, r *http.Request) bool {
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
		if claims.Role == "none" {
			h.writeUnauthorized(w)
			return false
		}

		// validate if session is active

		if claims.SessionId == 0 {
			h.writeUnauthorized(w)
			return false
		}

		expired := h.SessionGateway.SessionIsExpired(claims.SessionId)
		if expired {
			h.writeUnauthorized(w)
			return false
		}

		return true
	}
}
