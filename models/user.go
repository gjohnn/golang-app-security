package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" validate:"required,min=3,max=50"`
	Password string `gorm:"size:255;not null" json:"-" validate:"required,min=8"`
	Email    string `gorm:"size:255;not null;unique" json:"email" validate:"required,email"`
	Role     string `gorm:"default:USER"`
}

// LoginRequest es la estructura para la solicitud de login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RegisterRequest es la estructura para la solicitud de registro
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
