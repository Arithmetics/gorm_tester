package models

import "github.com/jinzhu/gorm"

//Game struct
type Game struct {
	gorm.Model
	Name   string
	Active bool
}
