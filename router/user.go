package router

import (
	"net/http"
	"restService/internals/handler"
)

func RegisterUserRoute(mux *http.ServeMux, handler *handler.UserHandler) {
	userHandler := http.NewServeMux()

	userHandler.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong from user"))
	})

	userHandler.HandleFunc("/", handler.GetUser)

	mux.Handle("/users/", http.StripPrefix("/users", userHandler))
}
