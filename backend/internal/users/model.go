package users

import "gorm.io/gorm"

// User merepresentasikan entitas pengguna di sistem.
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
