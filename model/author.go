package model

import (
	"gorm.io/gorm"
)

type Author struct{
	gorm.Model
	Name string
	Age int
	Book []Book `gorm:"many2many:book_author;"`
}