package handler

import (
	"encoding/json"
	"net/http"
	"restService/internals/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := h.userService.GetAllUser()
	bytes, _ := json.Marshal(users)
	w.Write(bytes)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.userService.GetUser(1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	bytes, _ := json.Marshal(user)
	w.Write(bytes)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := h.userService.CreateUser()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("User created successfully"))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := h.userService.UpdateUser(1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("User updated successfully"))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := h.userService.DeleteUser()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("User deleted successfully"))
}
