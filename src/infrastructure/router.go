package infrastructure

import "net/http"

func SetupRoutes(pc *ProductController) {
	http.HandleFunc("/products", pc.CreateNewHandler)
	http.HandleFunc("/products/all", pc.GetAllHandler)
	http.HandleFunc("/products/update", pc.UpdateHandler)
	http.HandleFunc("/products/delete", pc.DeleteHandler)
}
