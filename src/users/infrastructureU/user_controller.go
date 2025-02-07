package infrastructureU

import (
	"ArquitecturaExagonal/src/users/applicationU"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	CreateUseCase *applicationU.CreateUser
	GetAllUseCase *applicationU.GetAllUsers
	UpdateUseCase *applicationU.UpdateUser
	DeleteUseCase *applicationU.DeleteUser
}

func NewUserController(
	create *applicationU.CreateUser,
	getAll *applicationU.GetAllUsers,
	update *applicationU.UpdateUser,
	delete *applicationU.DeleteUser,
) *UserController {
	return &UserController{
		CreateUseCase: create,
		GetAllUseCase: getAll,
		UpdateUseCase: update,
		DeleteUseCase: delete,
	}
}

func (uc *UserController) CreateNewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var userInput struct {
		Name  string `json:"name"`
		Phone string `json:"phone_number"`
	}
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Datos recibidos: %+v\n", userInput)

	err = uc.CreateUseCase.Run(userInput.Name, userInput.Phone)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear el usuario: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuario creado exitosamente"))
}

func (uc *UserController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	users, err := uc.GetAllUseCase.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuarios: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var userInput struct {
		ID    int32  `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone_number"`
	}
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}
	err = uc.UpdateUseCase.Run(userInput.ID, userInput.Name, userInput.Phone)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario actualizado"))
}

func (uc *UserController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var userInput struct {
		ID int32 `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer JSON: %v", err), http.StatusBadRequest)
		return
	}
	err = uc.DeleteUseCase.Run(userInput.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el usuario: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado correctamente"))
}
