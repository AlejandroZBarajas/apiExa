package application

import "ArquitecturaExagonal/src/domain"

type GetAllProducts struct {
	repo domain.ProductInterface
}

func NewGetAllProducts(repo domain.ProductInterface) *GetAllProducts {
	return &GetAllProducts{repo: repo}
}

func (gap *GetAllProducts) Run([]*domain.ProductInterface, error) {
	return gap.repo.GetAll()
}
