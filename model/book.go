package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookName string `json:"book_name" gorm:"column:book_name"`
	Author   string `json:"author" gorm:"column:author"`
	Qty      int32  `json:"qty" gorm:"column:qty"`
}
