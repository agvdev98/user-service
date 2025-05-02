package main

import (
	"github.com/agvdev98/user-service/initializers"
	"github.com/agvdev98/user-service/internal/db"
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
	//userRepo := repository.NewUserRepository(database)
	//userService := service.NewUserService(userRepo)

	// Router
	r := gin.Default()

	// Routes
	r.GET("/users", func(c *gin.Context) {
		// Aquí usarías userService.FindAllUsers()

	})

	// Run server
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
