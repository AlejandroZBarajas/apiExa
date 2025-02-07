package infrastructureP

import (
	"ArquitecturaExagonal/src/products/applicationP"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductController struct {
	CreateUseCase *applicationP.CreateProduct
	GetAllUseCase *applicationP.GetAllProducts
	UpdateUseCase *applicationP.UpdateProduct
	DeleteUseCase *applicationP.DeleteProduct
}

func NewProductController(
	create *applicationP.CreateProduct,
	getAll *applicationP.GetAllProducts,
	update *applicationP.UpdateProduct,
	delete *applicationP.DeleteProduct,
) *ProductController {
	return &ProductController{
		CreateUseCase: create,
		GetAllUseCase: getAll,
		UpdateUseCase: update,
		DeleteUseCase: delete,
	}
}

func (pc *ProductController) CreateNewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var ProductInput struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	err := json.NewDecoder(r.Body).Decode(&ProductInput)
	if err != nil {

		http.Error(w, fmt.Sprintf("error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err = pc.CreateUseCase.Run(ProductInput.Name, ProductInput.Price)
	if err != nil {
		http.Error(w, fmt.Sprintf("error al crear el producto : %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("producto creado exitosamente"))
}

func (pc *ProductController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	products, err := pc.GetAllUseCase.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener productos: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(products)
}

func (pc *ProductController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var productInput struct {
		Id    int32   `json:"id"`
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error : %v", err), http.StatusBadRequest)
		return
	}

	/* idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	} */

	if productInput.Id <= 0 {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	/* price, err := strconv.ParseFloat(productInput.Price, 32)
	if err != nil {
		http.Error(w, "Precio invalido", http.StatusBadRequest)
		return
	} */
	err = pc.UpdateUseCase.Run(productInput.Id, productInput.Name, productInput.Price)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar: %v", err), http.StatusInternalServerError)
		return
	}

	/* err = pc.UpdateUseCase.Run(int32(id), productInput.Name, float32(price))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	} */

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto actualizado"))
}

func (pc *ProductController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var productInput struct {
		ID int32 `json:"id"` // Recibe ID desde el body en JSON
	}

	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer JSON: %v", err), http.StatusBadRequest)
		return
	}

	if productInput.ID <= 0 {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	err = pc.DeleteUseCase.Run(productInput.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el producto: %v", err), http.StatusInternalServerError)
		return
	}
	/*
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID invalido", http.StatusBadRequest)
			return
		}

		err = pc.DeleteUseCase.Run(int32(id))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error al eliminar el producto: %v", err), http.StatusInternalServerError)
			return
		} */
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto eliminado correctamente"))
}
