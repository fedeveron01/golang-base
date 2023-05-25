package infrastructure

import (
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
	handler_person "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/person"
	"github.com/fedeveron01/golang-base/cmd/usecases/calculate_age"
)

//inject dependencies..

type HandlerContainer struct {
	CalculateAge entrypoints.Handler
}

func Start() HandlerContainer {

	// inject repositories
	// inject use cases
	calculateAgeUseCase := calculate_age.Implementation{}

	// inject handlers
	handlerContainer := HandlerContainer{}
	handlerContainer.CalculateAge = handler_person.NewPersonGetAllHandler(calculateAgeUseCase)

	//handlerContainer.CalculateAge = calculateAgeHandler
	return handlerContainer

}
