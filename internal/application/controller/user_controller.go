package controller

import (
	"encoding/json"
	"net/http"

	"github.com/NiltonMorais/event-driven-golang/internal/application/dto"
	"github.com/NiltonMorais/event-driven-golang/internal/application/usecase"
)

type UserController struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserController(createUserUseCase *usecase.CreateUserUseCase) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
	}
}

func (u *UserController) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestData dto.CreateUserDTO
	json.NewDecoder(r.Body).Decode(&requestData)
	err := u.createUserUseCase.Execute(r.Context(), requestData.Name, requestData.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
