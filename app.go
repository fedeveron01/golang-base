package main

import (
	"fmt"
	"time"

	"github.com/fedeveron01/golang-base/entities"
	person_gateway "github.com/fedeveron01/golang-base/entrypoints/gateways/person"
	"github.com/fedeveron01/golang-base/usecases/calculate_age"
)

func travelArray(array []string) []string {
	for _, element := range array {
		fmt.Println(element)
	}

	return array
}

func main() {
	useCaseImplementation := calculate_age.Implementation{}
	gatewayImplementation := person_gateway.Gateway{}

	person := entities.Person{Id: 1, Name: "fede", LastName: "veron", BornDate: time.Date(2000, 2, 19, 14, 0, 0, 0, time.Local)}

	useCasePerson := gatewayImplementation.MapEntitiePersonToPersonUseCase(person)
	res := useCaseImplementation.CalculateAge(useCasePerson)
	fmt.Print(res)
}
