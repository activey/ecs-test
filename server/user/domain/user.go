package domain

type User struct {
	Id       uint
	Name     string
	Password string
}

func NewUser(id uint, name string, password string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Password: password,
	}
}
