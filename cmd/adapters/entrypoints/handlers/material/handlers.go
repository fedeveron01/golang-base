package material_handler

import (
	"encoding/json"
	"fmt"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/usecases/material"
	"io"
	"net/http"
)

type MaterialHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}
type MaterialHandler struct {
	entrypoints.HandlerBase
	materialUseCase material_usecase.MaterialUseCase
}

func NewMaterialHandler(sessionGateway gateways.SessionGateway, materialUseCase material_usecase.MaterialUseCase) *MaterialHandler {
	return &MaterialHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		materialUseCase: materialUseCase,
	}
}

// Create api/material
func (p *MaterialHandler) Create(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	reqBody, _ := io.ReadAll(r.Body)
	var material entities.Material
	json.Unmarshal(reqBody, &material)
	err := p.materialUseCase.CreateMaterial(material)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}
	p.WriteResponse(w, "material created", http.StatusCreated)

}

func (p *MaterialHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	materials, err := p.materialUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(materials)
}
