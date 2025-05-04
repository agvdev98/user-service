package service

import (
	"errors"
	"github.com/agvdev98/user-service/internal/model"
	"github.com/agvdev98/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	if strings.TrimSpace(user.Email) == "" {
		return nil, errors.New("email is required")
	}

	if strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

func (s *userServiceImpl) FindUserByID(id uint) (*model.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *userServiceImpl) FindAllUsers() ([]model.User, error) {
	return s.repo.FindAllUsers()
}

func (s *userServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	existingUser, err := s.repo.FindUserByID(user.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if strings.TrimSpace(user.Email) != "" {
		existingUser.Email = user.Email
	}

	if strings.TrimSpace(user.Name) != "" {
		existingUser.Name = user.Name
	}

	if strings.TrimSpace(user.Password) != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		existingUser.Password = string(hashedPassword)
	}

	return s.repo.UpdateUser(existingUser)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	user, err := s.repo.FindUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repo.DeleteUser(user.ID)
}
