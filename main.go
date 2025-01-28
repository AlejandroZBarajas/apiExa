package main

import (
	"ArquitecturaExagonal/src/application"
	"ArquitecturaExagonal/src/infrastructure"
	"fmt"
	"log"
	"net/http"
)

func main() {
	infrastructure.ConnectDB()
	db := infrastructure.GetDB()                           //jala la db
	productRepo := infrastructure.NewProductRepository(db) //nuevo repo

	createUseCase := application.NewProductCreation(productRepo)
	getAllUseCase := application.NewGetAllProducts(productRepo)
	updateUseCase := application.NewUpdateProduct(productRepo)
	deleteUseCase := application.NewDeleteProduct(productRepo)

	productController := infrastructure.NewProductController(
		createUseCase,
		getAllUseCase,
		updateUseCase,
		deleteUseCase,
	)

	infrastructure.SetupRoutes(productController)

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
