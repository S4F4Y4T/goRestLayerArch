package handler

import (
	"encoding/json"
	"net/http"
	"restService/internals/model"
	"restService/internals/service"
	"restService/pkg"
	"restService/pkg/validator"
	"strconv"
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
	users, err := h.userService.GetAllUser()
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	pkg.SuccessResponse(w, http.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		pkg.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	pkg.SuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validation
	if validator.IsEmpty(user.Name) {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if !validator.ValidateEmail(user.Email) {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	if err := h.userService.CreateUser(&user); err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.SuccessResponse(w, http.StatusCreated, "User created successfully", user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Check if user exists
	existingUser, err := h.userService.GetUser(uint(id))
	if err != nil {
		pkg.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	var updateData model.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validation
	if validator.IsEmpty(updateData.Name) {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if !validator.ValidateEmail(updateData.Email) {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	// Update fields
	existingUser.Name = updateData.Name
	existingUser.Email = updateData.Email

	if err := h.userService.UpdateUser(existingUser); err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.SuccessResponse(w, http.StatusOK, "User updated successfully", existingUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Check if user exists
	existingUser, err := h.userService.GetUser(uint(id))
	if err != nil {
		pkg.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if err := h.userService.DeleteUser(existingUser); err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.SuccessResponse(w, http.StatusOK, "User deleted successfully", nil)
}
