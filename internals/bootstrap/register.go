package bootstrap

import (
	"restService/internals/handler"
	"restService/internals/service"
)

type App struct {
	UserHandler *handler.UserHandler
}

func Register() *App {
	service := service.NewUserService()
	handler := handler.NewUserHandler(service)

	return &App{
		UserHandler: handler,
	}
}
