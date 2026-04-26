package service

import (
	"restService/internals/model"
)

type UserService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUser() ([]model.User, error) {
	return s.repo.FindAll()
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
