package user_handler

type LoginRequest struct {
	UserName string
	Password string
}

type CreateUserRequest struct {
	UserName string
	Password string
	Name     string
	LastName string
	DNI      string
	Charge   string
}

type TokenResponse struct {
	Token      string
	EmployeeId float64
	Charge     string
}
