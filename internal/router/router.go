package router

import (
	"github.com/gin-gonic/gin"
	"github.com/agvdev98/user-service/internal/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	// Routes
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)

	}
	return router
}
