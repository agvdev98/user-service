package repository

import "github.com/agvdev98/user-service/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	FindUserByID(id uint) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id uint) error
}
