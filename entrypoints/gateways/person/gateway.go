package person_gateway

import (
	"github.com/fedeveron01/golang-base/entities"
	"github.com/fedeveron01/golang-base/usecases/calculate_age"
)

type Gateway struct {
}

func (g *Gateway) MapEntitiePersonToPersonUseCase(person entities.Person) calculate_age.Person {
	return calculate_age.Person{
		Id:       person.Id,
		Name:     person.Name,
		LastName: person.LastName,
		BornDate: person.BornDate,
	}
}
