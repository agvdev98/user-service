package main

import (
	"github.com/agvdev98/user-service/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	router := gin.Default()

	router.Run(":8080")
}
