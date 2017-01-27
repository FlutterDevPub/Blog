package mocking

type UserStore interface {
	GetUser(name string) (*User, error)
	SetUser(name string, user *User) error
}

type User struct {
	Name    string
	Surname string
	Age     int
}
