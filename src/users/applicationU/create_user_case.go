package applicationU

import (
	"ArquitecturaExagonal/src/users/domainU"
	"ArquitecturaExagonal/src/users/domainU/userEntity"
	"fmt"
)

type CreateUser struct {
	repo domainU.UserInterface
}

func NewUserCreation(repo domainU.UserInterface) *CreateUser {
	return &CreateUser{repo: repo}
}

func (cu *CreateUser) Run(Name, Phone string) error {
	user := userEntity.CreateUser(Name, Phone)
	err := cu.repo.Save(user)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	return nil
}
