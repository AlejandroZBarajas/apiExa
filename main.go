package main

import (
	"ArquitecturaExagonal/src/core/infrastructureC"
	"ArquitecturaExagonal/src/products/applicationP"
	"ArquitecturaExagonal/src/products/infrastructureP"
	"ArquitecturaExagonal/src/users/applicationU"
	"ArquitecturaExagonal/src/users/infrastructureU"
	"fmt"
	"log"
	"net/http"
)

func main() {
	infrastructureC.ConnectDB()
	db := infrastructureC.GetDB()                           //jala la db
	productRepo := infrastructureP.NewProductRepository(db) //nuevo repo
	userRepo := infrastructureU.NewUserRepository(db)

	createUseCase := applicationP.NewProductCreation(productRepo)
	getAllUseCase := applicationP.NewGetAllProducts(productRepo)
	updateUseCase := applicationP.NewUpdateProduct(productRepo)
	deleteUseCase := applicationP.NewDeleteProduct(productRepo)

	createUserUseCase := applicationU.NewUserCreation(userRepo)
	getAllUserUseCase := applicationU.NewGetAllUsers(userRepo)
	updateUserUseCase := applicationU.NewUpdateUser(userRepo)
	deleteUserUseCase := applicationU.NewDeleteUser(userRepo)

	productController := infrastructureP.NewProductController(
		createUseCase,
		getAllUseCase,
		updateUseCase,
		deleteUseCase,
	)

	userController := infrastructureU.NewUserController(
		createUserUseCase,
		getAllUserUseCase,
		updateUserUseCase,
		deleteUserUseCase,
	)

	infrastructureC.SetupRoutes(productController, userController)

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
