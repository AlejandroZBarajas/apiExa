package applicationU

import "ArquitecturaExagonal/src/users/domainU"

//"fmt"

type UpdateUser struct {
	repo domainU.UserInterface
}

func NewUpdateUser(repo domainU.UserInterface) *UpdateUser {
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
