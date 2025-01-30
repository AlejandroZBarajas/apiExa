package domain

import "ArquitecturaExagonal/src/domain/entities"

type UserInterface interface {
	Save(user *entities.User) error
	GetAll() ([]*entities.User, error)
	Update(id int32, user *entities.User) error
	Delete(id int32) error
	GetByID(id int32) (*entities.User, error)
}
