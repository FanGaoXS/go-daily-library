package model

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{name, age}
}
