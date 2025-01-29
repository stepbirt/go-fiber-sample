package service

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type RequestNewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService interface {
	NewUser(RequestNewUser) (*UserResponse, error)
}
