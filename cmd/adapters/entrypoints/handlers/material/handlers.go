package material_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	material_usecase "github.com/fedeveron01/golang-base/cmd/usecases/material"
	"github.com/gorilla/mux"
)

type MaterialHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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

// Get api/material

func (p *MaterialHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	language := r.Header.Get("Language")
	if language == "" {
		language = "en"
	}
	materials, err := p.materialUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(ToMaterialsResponse(materials, language))
}

// Create api/material
func (p *MaterialHandler) Create(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	language := r.Header.Get("Language")
	if language == "" {
		language = "en"
	}
	reqBody, _ := io.ReadAll(r.Body)
	var materialRequest MaterialRequest
	json.Unmarshal(reqBody, &materialRequest)
	material, err := p.materialUseCase.CreateMaterial(ToMaterialEntity(materialRequest))
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	response := ToMaterialResponse(material, language)
	json.NewEncoder(w).Encode(response)

}

// Handle api/material/{id}
func (p *MaterialHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	if !p.IsAdmin(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		p.WriteErrorResponse(w, core_errors.NewBadRequestError("id is not valid"))
		return
	}
	err = p.materialUseCase.DeleteMaterial(id)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}
	p.WriteResponse(w, "material deleted", http.StatusOK)
	w.WriteHeader(http.StatusOK)
}
