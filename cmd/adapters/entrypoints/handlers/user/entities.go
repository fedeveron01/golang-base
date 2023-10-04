package user_handler

type LoginRequest struct {
	UserName string
	Password string
}

type CreateUserRequest struct {
	UserName string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
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
