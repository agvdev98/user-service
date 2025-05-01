package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:100" json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"-"` // Excluded from JSON responses (-)
}
