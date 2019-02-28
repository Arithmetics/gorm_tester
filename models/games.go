package models

import "github.com/jinzhu/gorm"

//Game struct
type Game struct {
	gorm.Model
	UserCreator uint
	Players     []User `gorm:"many2many:joined_games"`
	Name        string
	Active      bool
}
