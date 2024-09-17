package model

import (
	"gorm.io/gorm"
)

type Book struct{
	gorm.Model
	Title string `json:"title"`
	Price int `json:"price"`
	Author []Author `json:"author" gorm:"many2many:book_author;"`

}