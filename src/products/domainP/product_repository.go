package domainP

import "ArquitecturaExagonal/src/products/domainP/productEntity"

type ProductInterface interface {
	Save(product *productEntity.Product) error

	GetAll() ([]*productEntity.Product, error)

	Update(id int32, product *productEntity.Product) error

	Delete(id int32) error

	GetByID(id int32) (*productEntity.Product, error)
}
