package user

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ServetDeveloper/order-management/service/auth"
	"github.com/ServetDeveloper/order-management/types"
	"github.com/ServetDeveloper/order-management/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("burdayam1"))
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		var error1 validator.ValidationErrors
		errors.As(err, &error1)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("burdayam2"))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("burdayam3"))
	}

	// check password
	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ala parol"))
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": ""})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		var errorss validator.ValidationErrors
		errors.As(err, &errorss)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errorss))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJson(w, http.StatusCreated, nil)
	if err != nil {
		log.Fatal(err)
	}
}
