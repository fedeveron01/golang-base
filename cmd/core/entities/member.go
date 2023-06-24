package entities

type Member struct {
	Name        string `json:"name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
