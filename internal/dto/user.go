package dto

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
