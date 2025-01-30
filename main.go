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
	userRepo := infrastructure.NewUserRepository(db)

	createUseCase := application.NewProductCreation(productRepo)
	getAllUseCase := application.NewGetAllProducts(productRepo)
	updateUseCase := application.NewUpdateProduct(productRepo)
	deleteUseCase := application.NewDeleteProduct(productRepo)

	createUserUseCase := application.NewUserCreation(userRepo)
	getAllUserUseCase := application.NewGetAllUsers(userRepo)
	updateUserUseCase := application.NewUpdateUser(userRepo)
	deleteUserUseCase := application.NewDeleteUser(userRepo)

	productController := infrastructure.NewProductController(
		createUseCase,
		getAllUseCase,
		updateUseCase,
		deleteUseCase,
	)

	userController := infrastructure.NewUserController(
		createUserUseCase,
		getAllUserUseCase,
		updateUserUseCase,
		deleteUserUseCase,
	)

	infrastructure.SetupRoutes(productController, userController)

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
