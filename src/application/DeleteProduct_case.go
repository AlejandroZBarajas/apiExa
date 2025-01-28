package application

import "ArquitecturaExagonal/src/domain"

type DeleteProduct struct {
	repo domain.ProductInterface
}

func NewDeleteProduct(repo domain.ProductInterface) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (dp *DeleteProduct) Run(id int32) error {
	return dp.repo.Delete(id)
}
