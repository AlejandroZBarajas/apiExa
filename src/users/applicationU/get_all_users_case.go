package applicationU

import (
	"ArquitecturaExagonal/src/users/domainU"
	"ArquitecturaExagonal/src/users/domainU/userEntity"
)

type GetAllUsers struct {
	repo domainU.UserInterface
}

func NewGetAllUsers(repo domainU.UserInterface) *GetAllUsers {
	return &GetAllUsers{repo: repo}
}

func (gau *GetAllUsers) Run() ([]*userEntity.User, error) {
	return gau.repo.GetAll()
}
