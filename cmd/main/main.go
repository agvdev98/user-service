package main

import (
	"github.com/agvdev98/user-service/initializers"
	"github.com/agvdev98/user-service/internal/db"
	"github.com/agvdev98/user-service/internal/handler"
	"github.com/agvdev98/user-service/internal/repository"
	"github.com/agvdev98/user-service/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// Connect database
	database := db.Connect()

	// Sync database models
	initializers.SyncDatabase(database)

	// DI
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)

	// Init handler
	userHandler := handler.NewUserHandler(userService)

	// Router
	r := gin.Default()

	// Routes
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/user/:id", userHandler.GetUserByID)
	r.POST("/", userHandler.CreateUser)

	// Run server
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
