package models

import (
	"api-app/api/security"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `gorm:"unique" json:"email"`
	Password  string `json:"password,omitempty"`
	Address string `json:"address"`
	Role int `json:"role"`
	IgnoreMe     int     `gorm:"-"` // ignore this field
	Phone string
}

func (u *User) BeforeSave() error {
	password, err := security.Hash(u.Password)
	if  err != nil {
		return err
	}

	u.Password = string(password)
	return nil
}