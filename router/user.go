package router

import (
	"net/http"
	"restService/internals/handler"
)

func RegisterUserRoute(mux *http.ServeMux, handler *handler.UserHandler) {
	userHandler := http.NewServeMux()

	userHandler.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong from user"))
	})

	userHandler.HandleFunc("GET /", handler.GetAllUser)
	userHandler.HandleFunc("GET /{id}", handler.GetUser)
	userHandler.HandleFunc("POST /", handler.CreateUser)
	userHandler.HandleFunc("PUT /{id}", handler.UpdateUser)
	userHandler.HandleFunc("DELETE /{id}", handler.DeleteUser)

	mux.Handle("/users/", http.StripPrefix("/users", userHandler))
}
