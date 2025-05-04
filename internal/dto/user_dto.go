package dto

type UserRequestDTO struct {
	Name     string `json:"name" binding:"required,min=1,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequestDTO struct {
	Name     *string `json:"name" binding:"min=1,max=100,omitempty"`
	Email    *string `json:"email" binding:"email,omitempty"`
	Password *string `json:"password" binding:"min=6,omitempty"`
}

type UserResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
