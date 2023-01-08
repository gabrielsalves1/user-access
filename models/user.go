package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Number  int       `json:"number"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
