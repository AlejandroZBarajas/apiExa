package application

import (
	"ArquitecturaExagonal/src/domain"
	"ArquitecturaExagonal/src/domain/entities"
	"fmt"
)

type CreateProduct struct {
	repo domain.ProductInterface
}

func NewProductCreation(repo domain.ProductInterface) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Run(name string, price float32) error {
	product := entities.CreateProduct(name, price)
	err := cp.repo.Save(product)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	return nil
}
