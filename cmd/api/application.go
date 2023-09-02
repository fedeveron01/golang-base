package app

import (
	"github.com/fedeveron01/golang-base/cmd/infrastructure"
	"github.com/gorilla/mux"
)

func Start() {

	//configure mappings
	handlers := infrastructure.Start()
	r := mux.NewRouter()

	ConfigureMappings(*r, handlers)

}
