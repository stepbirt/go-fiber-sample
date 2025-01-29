package repository

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

type UserRepository interface {
	Create(User) (*User, error)
	// GetUser(int) (User, error)
	// GetUsers() ([]User, error)
}
