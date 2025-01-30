package application

import (
	"ArquitecturaExagonal/src/domain"
	//"fmt"
)

type UpdateUser struct {
	repo domain.UserInterface
}

func NewUpdateUser(repo domain.UserInterface) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (up *UpdateUser) Run(id int32, name, phoneNumber string) error {
	user, err := up.repo.GetByID(id)
	if err != nil {
		return err
	}
	user.SetName(name)
	user.SetPhoneNumber(phoneNumber)
	return up.repo.Update(id, user)
}
