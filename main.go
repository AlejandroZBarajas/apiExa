package main

import (
	"ArquitecturaExagonal/src/application"
	"ArquitecturaExagonal/src/infrastructure"
)

func main() {
	infrastructure.ConnectDB()
	db := infrastructure.GetDB()                           //jala la db
	productRepo := infrastructure.NewProductRepository(db) //nuevo repo

	createUseCase := application.NewProductCreation(productRepo)
}
