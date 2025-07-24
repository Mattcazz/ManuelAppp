package user

import (
	"fmt"
	"net/http"

	"github.com/Mattcazz/ManuelAppp.git/service/auth"
	"github.com/Mattcazz/ManuelAppp.git/types"
	"github.com/Mattcazz/ManuelAppp.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("Post")
	router.HandleFunc("/register", h.handleRegister).Methods("Post")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var registerUserPayload types.RegisterUserPayload
	err := utils.ParseJSON(r, registerUserPayload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err = h.store.GetUserByEmail(registerUserPayload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", registerUserPayload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(registerUserPayload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.store.CreateUser(types.User{
		UserName: registerUserPayload.UserName,
		Email:    registerUserPayload.Email,
		Password: hashedPassword,
	})

	if err == nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
