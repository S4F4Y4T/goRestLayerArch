package router

import (
	"net/http"
	"restService/internals/bootstrap"
	"restService/internals/middleware"
)

func SetUp(app *bootstrap.App) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	RegisterUserRoute(mux, app.UserHandler)

	return middleware.Apply(middleware.Logger, middleware.Cors, middleware.Recover)(mux)
}
