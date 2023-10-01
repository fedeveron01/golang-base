package material_type_handler

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	material_type_usecase "github.com/fedeveron01/golang-base/cmd/usecases/material_type"
	"io"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type MaterialTypeHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetUnitsOfMeasurement(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type MaterialTypeHandler struct {
	entrypoints.HandlerBase
	materialTypeUseCase material_type_usecase.MaterialTypeUseCase
}

func NewMaterialTypeHandler(sessionGateway gateways.SessionGateway, materialTypeUseCase material_type_usecase.MaterialTypeUseCase) *MaterialTypeHandler {
	return &MaterialTypeHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		materialTypeUseCase: materialTypeUseCase,
	}
}

func (p *MaterialTypeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	materialTypes, err := p.materialTypeUseCase.FindAll()
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}
	language := r.Header.Get("Language")
	if language == "" {
		language = "en"
	}
	materials := ToMaterialTypesResponse(materialTypes, language)

	json.NewEncoder(w).Encode(materials)
}

func (p *MaterialTypeHandler) GetUnitsOfMeasurement(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	language := r.Header.Get("Language")
	if language == "" {
		language = "en"
	}
	unitsOfMeasurement := enums.GetAllUnitOfMeasurementEnum(language)
	json.NewEncoder(w).Encode(ToUnitsOfMeasurementResponse(unitsOfMeasurement))
}

// Handle api/materialType
func (p *MaterialTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	if !p.IsAdmin(w, r) {
		return
	}
	reqBody, _ := io.ReadAll(r.Body)
	var materialType entities.MaterialType
	json.Unmarshal(reqBody, &materialType)
	// validate enum
	if enums.StringToUnitOfMeasurementEnum(materialType.UnitOfMeasurement.String("en")) == ("") {
		p.WriteErrorResponse(w, core_errors.NewBadRequestError("unitOfMeasurement is not valid"))
		return
	}
	err := p.materialTypeUseCase.CreateMaterialType(materialType)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	p.WriteResponse(w, "materialType created", http.StatusCreated)
	w.WriteHeader(http.StatusCreated)

}

// Handle api/materialType
func (p *MaterialTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	if !p.IsAdmin(w, r) {
		return
	}

	language := r.Header.Get("Language")
	if language == "" {
		language = "en"
	}

	reqBody, _ := io.ReadAll(r.Body)
	var materialType entities.MaterialType
	json.Unmarshal(reqBody, &materialType)

	// validate enum
	if enums.StringToUnitOfMeasurementEnum(materialType.UnitOfMeasurement.String("")) == ("") {
		p.WriteErrorResponse(w, core_errors.NewBadRequestError("unitOfMeasurement is not valid"))
		return
	}
	materialType, err := p.materialTypeUseCase.UpdateMaterialType(materialType)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	json.NewEncoder(w).Encode(ToMaterialTypeResponse(materialType, language))
}
