package main

import "time"

type Person struct {
	Id       int
	Name     string
	LastName string
	BornDate time.Time
}

// getter and setters by default

func (p *Person) CalculateAge() int {
	return 1
}
