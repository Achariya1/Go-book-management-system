package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" db:"username"` 
	Password string `json:"password" db:"password"` 
}


type SignupRequest struct{
	Username string `json:"username"` 
	Password string `json:"password"`
}



















