package applicationU

import (
	"ArquitecturaExagonal/src/users/domainU"
	"ArquitecturaExagonal/src/users/domainU/userEntity"
)

type GetByName struct {
	repo domainU.UserInterface
}

func NewGetByName(repo domainU.UserInterface) *GetByName {
	return &GetByName{repo: repo}
}

func (gbnu *GetByName) Run(name string) (*userEntity.User, error) {
	user, err := gbnu.repo.GetByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
