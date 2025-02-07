package infrastructureC

import (
	"ArquitecturaExagonal/src/products/infrastructureP"
	"ArquitecturaExagonal/src/users/infrastructureU"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(pc *infrastructureP.ProductController, uc *infrastructureU.UserController) {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", pc.CreateNewHandler)
	mux.HandleFunc("/products/all", pc.GetAllHandler)
	mux.HandleFunc("/products/update", pc.UpdateHandler)
	mux.HandleFunc("/products/delete", pc.DeleteHandler)

	mux.HandleFunc("/users", uc.CreateNewHandler)
	mux.HandleFunc("/users/all", uc.GetAllHandler)
	mux.HandleFunc("/users/update", uc.UpdateHandler)
	mux.HandleFunc("/users/delete", uc.DeleteHandler)

	http.Handle("/", corsMiddleware(mux))
}

/* package infrastructure

import "net/http"

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func corsMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler(w, r)
	}
}

func SetupRoutes(pc *ProductController, uc *UserController) {
	http.HandleFunc("/products", corsMiddleware(pc.CreateNewHandler))
	http.HandleFunc("/products/all", corsMiddleware(pc.GetAllHandler))
	http.HandleFunc("/products/update", corsMiddleware(pc.UpdateHandler))
	http.HandleFunc("/products/delete", corsMiddleware(pc.DeleteHandler))

	http.HandleFunc("/users", corsMiddleware(uc.CreateNewHandler))
	http.HandleFunc("/users/all", corsMiddleware(uc.GetAllHandler))
	http.HandleFunc("/users/update", corsMiddleware(uc.UpdateHandler))
	http.HandleFunc("/users/delete", corsMiddleware(uc.DeleteHandler))
} */

/* func SetupRoutes(pc *ProductController, uc *UserController) {
	http.HandleFunc("/products", pc.CreateNewHandler)
	http.HandleFunc("/products/all", pc.GetAllHandler)
	http.HandleFunc("/products/update", pc.UpdateHandler)
	http.HandleFunc("/products/delete", pc.DeleteHandler)

	http.HandleFunc("/users", uc.CreateNewHandler)
	http.HandleFunc("/users/all", uc.GetAllHandler)
	http.HandleFunc("/users/update", uc.UpdateHandler)
	http.HandleFunc("/users/delete", uc.DeleteHandler)
}
*/
