package models

import "gorm.io/gorm"

type System struct {
	gorm.Model
	Name  string  `json:"name"`
	Url   string  `json:"url"`
	Users []*User `gorm:"many2many:user_systems;"`
}
