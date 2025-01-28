package domain

import "ArquitecturaExagonal/src/domain/entities"

type ProductInterface interface {
	Save(product *entities.Product) error

	GetAll() ([]*entities.Product, error)

	Update(id int32, product *entities.Product) error

	Delete(id int32) error

	GetByID(id int32) (*entities.Product, error)
}
