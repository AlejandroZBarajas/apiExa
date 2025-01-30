package application

import (
	"ArquitecturaExagonal/src/domain"
	"ArquitecturaExagonal/src/domain/entities"
	"fmt"
)

type CreateUser struct {
	repo domain.UserInterface
}

func NewUserCreation(repo domain.UserInterface) *CreateUser {
	return &CreateUser{repo: repo}
}

func (cu *CreateUser) Run(name, phoneNumber string) error {
	user := entities.CreateUser(name, phoneNumber)
	err := cu.repo.Save(user)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	return nil
}
