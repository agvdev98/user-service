package model

//User represents a user entity in the system.
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`     // Primary key
	Name     string `gorm:"size:100" json:"name"`     // Max 100 characters
	Email    string `gorm:"uniqueIndex" json:"email"` // Unique constraint
	Password string `gorm:"not null" json:"-"`        // Excluded from JSON responses (-)
}
