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
