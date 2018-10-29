package user

import "log"

type UserUsecase interface {
	View(id int) (*User, error)
	Save(name string, age int) error
}

type userUsecase struct {
	repository UserRepository
}

func NewUserUsecase(ur UserRepository) *userUsecase {
	return &userUsecase{repository: ur}
}

func (uu *userUsecase) View(id int) (*User, error) {
	user, err := uu.repository.View(id)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (uu *userUsecase) Store(name string, age int) error {
	return nil
}
