package service

import (
	"github.com/agvdev98/user-service/internal/model"
	"github.com/agvdev98/user-service/internal/repository"
)

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *userServiceImpl) FindUserByID(id uint) (*model.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *userServiceImpl) FindAllUsers() ([]model.User, error) {
	return s.repo.FindAllUsers()
}

func (s *userServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
