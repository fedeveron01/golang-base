package main

import (
	"fmt"
	"time"
)

func recorrerLista(lista []string) []string {
	for _, element := range lista {
		fmt.Println(element)
	}

	return lista
}

func main() {
	a := 2
	b := 3
	c := a * b

	if a < b && b <= 2 {
		fmt.Println("a es menor a b y b es menor igual a 2")
	} else {
		fmt.Println("ok", c)
	}

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("otro")
	}

	lista := []string{
		"hola", "chau",
	}

	persona := Person{1, "fede", "fede", time.Now()}
	fmt.Println("edad", persona.CalculateAge())

	recorrerLista(lista)

}
