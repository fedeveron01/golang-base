package employee_handler

type EmployeeResponse struct {
	Id       uint           `json:"id"`
	Name     string         `json:"name"`
	LastName string         `json:"lastName"`
	DNI      string         `json:"dni"`
	Charge   ChargeResponse `json:"charge"`
	User     UserResponse   `json:"user"`
}
type UserResponse struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	Inactive bool   `json:"inactive"`
}

type ChargeResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
