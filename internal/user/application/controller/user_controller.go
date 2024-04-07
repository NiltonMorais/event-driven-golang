package controller

import (
	"encoding/json"
	"net/http"

	"github.com/NiltonMorais/event-driven-golang/internal/user/application/dto"
	"github.com/NiltonMorais/event-driven-golang/internal/user/application/usecase"
	"github.com/NiltonMorais/event-driven-golang/internal/user/domain/event"
)

type UserController struct {
	createUserUseCase       *usecase.CreateUserUseCase
	sendWelcomeEmailUseCase *usecase.SendWelcomeEmailUseCase
}

func NewUserController(createUserUseCase *usecase.CreateUserUseCase, sendWelcomeEmailUseCase *usecase.SendWelcomeEmailUseCase) *UserController {
	return &UserController{
		createUserUseCase:       createUserUseCase,
		sendWelcomeEmailUseCase: sendWelcomeEmailUseCase,
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

func (u *UserController) SendWelcomeEmail(w http.ResponseWriter, r *http.Request) {
	var body event.UserRegisteredEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := u.sendWelcomeEmailUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
