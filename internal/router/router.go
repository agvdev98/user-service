package router

import (
	"github.com/agvdev98/user-service/internal/handler"
	"github.com/agvdev98/user-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler, authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()

	// Routes
	protected := router.Group("/users")
	protected.Use(middleware.JWTMiddleware(), middleware.RequireRole("ADMIN"))
	{
		protected.GET("/all", userHandler.GetAllUsers)
		protected.GET("/:id", userHandler.GetUserByID)
		protected.PUT("/:id", userHandler.UpdateUser)
		protected.DELETE("/:id", userHandler.DeleteUser)

	}

	public := router.Group("/users")
	public.Use(middleware.JWTMiddleware())
	{
		public.POST("/register", userHandler.CreateUser)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}

	return router
}
