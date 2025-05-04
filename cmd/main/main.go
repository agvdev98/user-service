package main

import (
	"github.com/agvdev98/user-service/initializers"
	"github.com/agvdev98/user-service/internal/db"
	"github.com/agvdev98/user-service/internal/handler"
	"github.com/agvdev98/user-service/internal/repository"
	"github.com/agvdev98/user-service/internal/router"
	"github.com/agvdev98/user-service/internal/service"
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
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler)

	// Run server
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
