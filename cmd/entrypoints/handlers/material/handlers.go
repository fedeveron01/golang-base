package material_handler

import (
	"encoding/json"
	"fmt"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/usecases/material"
	"io"
	"net/http"
)

type CreateMaterialHandler struct {
	materialUsecase material_usecase.Implementation
}

type GetAllMaterialHandler struct {
	materialUsecase material_usecase.Implementation
}

func NewCreateMaterialHandler(materialUsecase material_usecase.Implementation) CreateMaterialHandler {
	return CreateMaterialHandler{
		materialUsecase: materialUsecase,
	}
}

func (p CreateMaterialHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var material entities.Material
	json.Unmarshal(reqBody, &material)
	err := p.materialUsecase.CreateMaterial(material)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
}

func NewGetAllMaterialHandler(materialUsecase material_usecase.Implementation) GetAllMaterialHandler {
	return GetAllMaterialHandler{
		materialUsecase: materialUsecase,
	}
}

func (p GetAllMaterialHandler) Handle(w http.ResponseWriter, r *http.Request) {
	materials, err := p.materialUsecase.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(materials)
}
