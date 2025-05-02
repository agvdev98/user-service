package initializers

import (
	"github.com/agvdev98/user-service/internal/model"
	"gorm.io/gorm"
	"log"
)

func SyncDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}
	log.Println("Database migrated successfully")
}
