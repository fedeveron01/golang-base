package handler_person

import (
	"fmt"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/usecases/calculate_age"
)

//get all

type PersonGetAllHandler struct {
	// use cases
	CalculateAge calculate_age.CalculateAgeUseCase
}

func NewPersonGetAllHandler(calculateAge calculate_age.CalculateAgeUseCase) PersonGetAllHandler {
	return PersonGetAllHandler{
		CalculateAge: calculateAge,
	}
}

func (p PersonGetAllHandler) Handle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("PERSON GET ALL HANDLER")
	//p.CalculateAge.CalculateAge()

}
