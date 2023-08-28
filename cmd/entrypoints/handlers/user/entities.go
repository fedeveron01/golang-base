package user_handler

type LoginRequest struct {
	UserName string
	Password string
}

type TokenResponse struct {
	Token string
}
