package model

type User struct {
	ID       uint   `gorm:"primaryKey" `
	Name     string `gorm:"size:100"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"not null" json:"-"`
}
