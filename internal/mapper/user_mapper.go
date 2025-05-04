package mapper

import (
	"github.com/agvdev98/user-service/internal/dto"
	"github.com/agvdev98/user-service/internal/model"
)

func ToUser(dto dto.UserRequestDTO) *model.User {
	return &model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

}

func ToUserDTO(user *model.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUpdatedUser(id uint, dto dto.UserUpdateRequestDTO) *model.User {
	user := &model.User{
		ID: id,
	}

	if dto.Name != nil {
		user.Name = *dto.Name
	}
	if dto.Email != nil {
		user.Email = *dto.Email
	}
	if dto.Password != nil {
		user.Password = *dto.Password
	}

	return user
}

func ToUserDTOList(users []model.User) []dto.UserResponseDTO {
	res := make([]dto.UserResponseDTO, len(users))
	for i, u := range users {
		res[i] = ToUserDTO(&u)
	}
	return res
}
