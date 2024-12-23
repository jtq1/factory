package models

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Users []User

var DefaultUsers = Users{
	{
		ID:       1,
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
	},
}
