package user_handler

type LoginRequest struct {
	UserName string
	Password string
}

type ActiveDesactiveUserRequest struct {
	Inactive bool `json:"inactive" required:"true"`
}

type CreateUserRequest struct {
	UserName string
	Password string
	Name     string
	LastName string
	DNI      string
	ChargeId float64
}

type TokenResponse struct {
	Token      string
	EmployeeId float64
	Charge     string
}
