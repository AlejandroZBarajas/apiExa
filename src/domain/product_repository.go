package domain

import "ArquitecturaExagonal/src/domain/entities"

type ProductInterface interface {
	Save(product *entities.Product) error

	GetAll() ([]*entities.Product, error)

	Update(id int32, name string, price float32) error

	Delete(id int32) error
}
