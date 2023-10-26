package product_handler

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	product_usecase "github.com/fedeveron01/golang-base/cmd/usecases/product"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"

	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type ProductHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	AssignMaterialsToProduct(w http.ResponseWriter, r *http.Request)
}

type ProductHandler struct {
	entrypoints.HandlerBase
	productUseCase product_usecase.ProductUseCase
}

func NewProductHandler(sessionGateway gateways.SessionGateway, productUseCase product_usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		productUseCase: productUseCase,
	}
}

// GetAll Handle api/product GET request
func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	products, err := p.productUseCase.FindAll()
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	productsResponse := ToProductsResponse(products)

	json.NewEncoder(w).Encode(productsResponse)
}

// GetById Handle api/product/{id} GET request
func (p *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.WriteErrorResponse(w, core_errors.NewBadRequestError("id is not valid"))
		return
	}
	product, err := p.productUseCase.FindById(uint(uid))
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	language := r.Header.Get("Language")
	if language == "" {
		language = "en"
	}

	productResponse := ToProductWithAssignationsResponse(product, language)
	json.NewEncoder(w).Encode(productResponse)
}

// Create Handle api/product POST request
func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	var product entities.Product
	json.Unmarshal(reqBody, &product)

	var err error
	product, err = p.productUseCase.CreateProduct(product)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	productResponse := ToProductResponse(product)
	json.NewEncoder(w).Encode(productResponse)

}

// Update Handle api/product PUT request
func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	var product entities.Product
	json.Unmarshal(reqBody, &product)

	var err error
	product, err = p.productUseCase.UpdateProduct(product)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	productResponse := ToProductResponse(product)
	json.NewEncoder(w).Encode(productResponse)

}

// Delete Handle api/product DELETE request
func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	if !p.IsAdmin(w, r) {
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.WriteErrorResponse(w, core_errors.NewBadRequestError("id is not valid"))
		return
	}

	err = p.productUseCase.DeleteProduct(uint(uid))
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// AssignMaterialsToProduct Handle api/product/assignation POST request
func (p *ProductHandler) AssignMaterialsToProduct(w http.ResponseWriter, r *http.Request) {
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
	var assignation MaterialsProductAssignationRequest
	json.Unmarshal(reqBody, &assignation)

	var err error
	materialProducts, err := p.productUseCase.AssignMaterialsToProduct(assignation.ProductId, ToMaterialsProductEntity(assignation.Assignations))
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}

	assignationResponse := ToAssignationsResponse(materialProducts, language)
	json.NewEncoder(w).Encode(assignationResponse)
}
