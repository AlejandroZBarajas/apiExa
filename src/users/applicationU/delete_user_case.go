package applicationU

import "ArquitecturaExagonal/src/users/domainU"

type DeleteUser struct {
	repo domainU.UserInterface
}

func NewDeleteUser(repo domainU.UserInterface) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (du *DeleteUser) Run(id int32) error {
	return du.repo.Delete(id)
}
