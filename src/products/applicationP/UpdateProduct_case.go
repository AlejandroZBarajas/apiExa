package applicationP

import (
	"ArquitecturaExagonal/src/products/domainP"
)

type UpdateProduct struct {
	repo domainP.ProductInterface
}

func NewUpdateProduct(repo domainP.ProductInterface) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (up *UpdateProduct) Run(id int32, name string, price float32) error {
	product, err := up.repo.GetByID(id)
	if err != nil {
		return err
	}
	product.SetName(name)
	product.SetPrice(price)

	return up.repo.Update(id, product)
}
