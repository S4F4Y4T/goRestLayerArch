package service

import (
	"restService/internals/model"
	"restService/internals/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUser() []model.User {
	users, _ := s.repo.FindAll()
	return users
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(user *model.User) error {
	return s.repo.Delete(user)
}
