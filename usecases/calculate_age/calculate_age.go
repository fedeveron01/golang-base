package calculate_age

type CalculateAgeUseCase interface {
	CalculateAge(person Person) int
}

type Implementation struct {
}

func (c *Implementation) CalculateAge(person Person) int {
	clock := ClockImplementation{}
	return person.CalculateAge(clock)
}
