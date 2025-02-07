package applicationP

import (
	"ArquitecturaExagonal/src/products/domainP"
	"ArquitecturaExagonal/src/products/domainP/productEntity"
	"fmt"
)

type CreateProduct struct {
	repo domainP.ProductInterface
}

func NewProductCreation(repo domainP.ProductInterface) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Run(name string, price float32) error {
	product := productEntity.CreateProduct(name, price)
	err := cp.repo.Save(product)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	return nil
}
