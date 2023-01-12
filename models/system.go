package models

import "gorm.io/gorm"

type System struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}
