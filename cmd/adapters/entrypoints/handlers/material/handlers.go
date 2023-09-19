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

type CreateMaterialHandler struct {
	entrypoints.HandlerBase
	materialUseCase material_usecase.MaterialUseCase
}

type GetAllMaterialHandler struct {
	entrypoints.HandlerBase
	materialUseCase material_usecase.MaterialUseCase
}

func NewCreateMaterialHandler(sessionGateway gateways.SessionGateway, materialUseCase material_usecase.MaterialUseCase) CreateMaterialHandler {
	return CreateMaterialHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		materialUseCase: materialUseCase,
	}
}

func (p CreateMaterialHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	reqBody, _ := io.ReadAll(r.Body)
	var material entities.Material
	json.Unmarshal(reqBody, &material)
	err := p.materialUseCase.CreateMaterial(material)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func NewGetAllMaterialHandler(sessionGateway gateways.SessionGateway, materialUseCase material_usecase.MaterialUseCase) GetAllMaterialHandler {
	return GetAllMaterialHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		materialUseCase: materialUseCase,
	}
}

func (p GetAllMaterialHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	materials, err := p.materialUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(materials)
}
