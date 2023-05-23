package calculate_age

import "time"

type Person struct {
	Id       int
	Name     string
	LastName string
	BornDate time.Time
}

// getter and setters by default

func (p *Person) CalculateAge(clock Clock) int {
	return clock.Now().Year() - p.BornDate.Year()
}

type Clock struct {
}

func (c *Clock) Now() time.Time {
	return time.Now()
}
