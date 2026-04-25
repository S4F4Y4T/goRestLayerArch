package bootstrap

import (
	"restService/internals/handler"
	"restService/internals/repository"
	"restService/internals/service"

	"gorm.io/gorm"
)

type App struct {
	UserHandler *handler.UserHandler
}

func Register(db *gorm.DB) *App {

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	return &App{
		UserHandler: handler,
	}
}
