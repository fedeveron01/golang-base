package handler_person

import (
	"fmt"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/core/usecases/calculate_age"

	_ "github.com/mattn/go-sqlite3"
)

//get all

type PersonGetAllHandler struct {
	// use cases
	CalculateAge       calculate_age.CalculateAgeUseCase
	WhatsappRepository repositories.WhatsappRepository
}

func NewPersonGetAllHandler(calculateAge calculate_age.CalculateAgeUseCase, whatsappRepository repositories.WhatsappRepository) PersonGetAllHandler {
	return PersonGetAllHandler{
		CalculateAge:       calculateAge,
		WhatsappRepository: whatsappRepository,
	}
}

func (p PersonGetAllHandler) Handle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phone := params["phone"]
	message := params["message"]

	p.WhatsappRepository.SendText(phone, message)
	fmt.Fprint(w, "Sending message "+message+" to "+phone)

	//p.CalculateAge.CalculateAge()

}
