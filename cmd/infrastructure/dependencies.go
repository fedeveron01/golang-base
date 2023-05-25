package infrastructure

import (
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
)

//inject dependencies..

type HandlerContainer struct {
	CalculateAge entrypoints.Handler
}

func Start() HandlerContainer {

	// inject repositories
	// inject use cases
	// inject handlers
	handlerContainer := HandlerContainer{}
	//handlerContainer.CalculateAge = calculateAgeHandler
	return handlerContainer

}
