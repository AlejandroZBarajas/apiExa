package domainU

import entities "ArquitecturaExagonal/src/users/domainU/userEntity"

type UserInterface interface {
	Save(user *entities.User) error
	GetAll() ([]*entities.User, error)
	Update(id int32, user *entities.User) error
	Delete(id int32) error
	GetByID(id int32) (*entities.User, error)
	GetByName(name string) (*entities.User, error)
}
