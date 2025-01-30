package infrastructure

import "net/http"

func SetupRoutes(pc *ProductController, uc *UserController) {
	http.HandleFunc("/products", pc.CreateNewHandler)
	http.HandleFunc("/products/all", pc.GetAllHandler)
	http.HandleFunc("/products/update", pc.UpdateHandler)
	http.HandleFunc("/products/delete", pc.DeleteHandler)

	http.HandleFunc("/users", uc.CreateNewHandler)
	http.HandleFunc("/users/all", uc.GetAllHandler)
	http.HandleFunc("/users/update", uc.UpdateHandler)
	http.HandleFunc("/users/delete", uc.DeleteHandler)
}
