package api

import (
	"fmt"
	"time"

	"github.com/fedeveron01/golang-base/cmd/entities"
	"github.com/fedeveron01/golang-base/cmd/infrastructure"
	"github.com/fedeveron01/golang-base/cmd/usecases/calculate_age"
	"github.com/gorilla/mux"
)

func Start() {
	useCaseImplementation := calculate_age.Implementation{}
	fmt.Println("OK")

	person := entities.Person{Id: 1, Name: "fede", LastName: "veron", BornDate: time.Date(2000, 2, 19, 14, 0, 0, 0, time.Local)}

	res := useCaseImplementation.CalculateAge(person)
	fmt.Println(res)

	//configure mappings
	handlers := infrastructure.Start()
	r := mux.NewRouter()

	ConfigureMappings(*r, handlers)

}
