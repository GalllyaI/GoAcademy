package models

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
}

func Login(email, password string) (*User, error) {

	return &User{}, nil
}
