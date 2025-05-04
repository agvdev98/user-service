package handler

import (
	"github.com/agvdev98/user-service/internal/dto"
	"github.com/agvdev98/user-service/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(service service.UserService) *AuthHandler {
	return &AuthHandler{userService: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var userLoginDTO dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&userLoginDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.userService.Authenticate(userLoginDTO.Email, userLoginDTO.Password)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
