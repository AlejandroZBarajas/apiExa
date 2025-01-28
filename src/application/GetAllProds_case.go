package application

import (
	"ArquitecturaExagonal/src/domain"
	"ArquitecturaExagonal/src/domain/entities"
)

type GetAllProducts struct {
	repo domain.ProductInterface
}

func NewGetAllProducts(repo domain.ProductInterface) *GetAllProducts {
	return &GetAllProducts{repo: repo}
}

func (gap *GetAllProducts) Run() ([]*entities.Product, error) {
	return gap.repo.GetAll()
}
