package dto

type UserRequestDTO struct {
	Name     string `json:"name" binding:"required,min=1,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequestDTO struct {
	Name     *string `json:"name" binding:"omitempty,min=1,max=100"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty,min=6"`
}

type UserResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
