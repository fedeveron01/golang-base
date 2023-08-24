package handler_person

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/core/usecases/calculate_age"

	_ "github.com/mattn/go-sqlite3"
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
	params := mux.Vars(r)
	phone := params["phone"]
	message := params["message"]

	fmt.Println(phone, message)

	//p.CalculateAge.CalculateAge()

}
