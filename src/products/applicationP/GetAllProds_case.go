package applicationP

import (
	"ArquitecturaExagonal/src/products/domainP"
	"ArquitecturaExagonal/src/products/domainP/productEntity"
)

type GetAllProducts struct {
	repo domainP.ProductInterface
}

func NewGetAllProducts(repo domainP.ProductInterface) *GetAllProducts {
	return &GetAllProducts{repo: repo}
}

func (gap *GetAllProducts) Run() ([]*productEntity.Product, error) {
	return gap.repo.GetAll()
}
