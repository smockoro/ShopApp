package user

import (
	"fmt"
)

type UserRepository interface {
	View(id int) (*User, error)
	Save(name string, age int) error
}

type userRepository struct {
	Conn string
}

func NewUserRepository(conn string) UserRepository {
	return &userRepository{Conn: conn}
}

func (ur *userRepository) View(id int) (*User, error) {
	user := &User{
		Name: "Adam",
		Age:  10000,
	}
	return user, nil
}

func (ur *userRepository) Save(name string, age int) error {
	fmt.Printf("%s : %d", name, age)
	return nil
}
