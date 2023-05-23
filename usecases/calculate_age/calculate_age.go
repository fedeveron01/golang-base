package calculate_age

type CalculateAgeUseCase interface {
	CalculateAge(person Person) int
}

type Implementation struct {
}

func CalculateAge(person Person) int {
	clock := Clock{}
	return person.CalculateAge(clock)
}
