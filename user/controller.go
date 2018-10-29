package user

import "log"

type UserController interface {
	View(id int) (*User, error)
	Save(name string, age int) error
}

type userController struct {
	Conn string
}

func NewUserController(conn string) UserController {
	return &userController{
		Conn: conn,
	}
}

func (uc *userController) View(id int) (*User, error) {
	// ロジック呼び出し
	rep := NewUserRepository(uc.Conn)
	uu := NewUserUsecase(rep)
	user, err := uu.View(id)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (uc *userController) Save(name string, age int) error {
	return nil
}
