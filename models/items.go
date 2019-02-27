package models

import "github.com/jinzhu/gorm"

//Item struct
type Item struct {
	gorm.Model
	Name   string
	Price  uint
	Action string
}
