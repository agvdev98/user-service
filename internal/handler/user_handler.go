package handler

import (
	"github.com/agvdev98/user-service/internal/dto"
	"log"
	"net/http"
	"strconv"

	"github.com/agvdev98/user-service/internal/mapper"
	"github.com/agvdev98/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

// Get /users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.ToUserDTOList(users))
}

// Get /users/:id
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userId := uint(id)

	user, err := h.userService.FindUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, mapper.ToUserDTO(user))
}

// Post
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userDTO dto.UserRequestDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	user := mapper.ToUser(userDTO)
	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, mapper.ToUserDTO(createdUser))
}

// Put /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userId := uint(id)

	var userDTO dto.UserRequestDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	user := mapper.ToUser(userDTO)
	user.ID = userId

	updatedUser, err := h.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, mapper.ToUserDTO(updatedUser))
}

// Delete /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userId := uint(id)

	err = h.userService.DeleteUser(userId)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
