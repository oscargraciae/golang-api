package models

import "github.com/jinzhu/gorm"

type Business struct {
	gorm.Model
	Name string
	Description string
	Summary string
	logo string
	isActive bool
}
