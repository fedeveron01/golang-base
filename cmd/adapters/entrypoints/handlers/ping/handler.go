package ping_handler

import (
	"encoding/json"
	"net/http"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ok")
}
