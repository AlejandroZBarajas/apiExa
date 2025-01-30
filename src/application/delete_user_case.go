package application

import "ArquitecturaExagonal/src/domain"

type DeleteUser struct {
	repo domain.UserInterface
}

func NewDeleteUser(repo domain.UserInterface) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (du *DeleteUser) Run(id int32) error {
	return du.repo.Delete(id)
}
