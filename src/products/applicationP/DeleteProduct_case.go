package applicationP

import (
	"ArquitecturaExagonal/src/products/domainP"
)

type DeleteProduct struct {
	repo domainP.ProductInterface
}

func NewDeleteProduct(repo domainP.ProductInterface) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (dp *DeleteProduct) Run(id int32) error {
	return dp.repo.Delete(id)
}
