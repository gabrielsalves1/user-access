package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	Country string    `json:"country"`
	Number  int       `json:"number"`
	Company string    `json:"company"`
	Team    string    `json:"team"`
	Systems []*System `gorm:"many2many:user_systems;"`
}
