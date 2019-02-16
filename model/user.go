package model

import "github.com/jinzhu/gorm"

// User is user model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
