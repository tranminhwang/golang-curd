package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"title"`
	Body  string `gorm:"body"`
}
