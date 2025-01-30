package application

import (
	"ArquitecturaExagonal/src/domain"
	"ArquitecturaExagonal/src/domain/entities"
)

type GetAllUsers struct {
	repo domain.UserInterface
}

func NewGetAllUsers(repo domain.UserInterface) *GetAllUsers {
	return &GetAllUsers{repo: repo}
}

func (gau *GetAllUsers) Run() ([]*entities.User, error) {
	return gau.repo.GetAll()
}
