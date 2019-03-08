package models

import "github.com/jinzhu/gorm"

//Faction (Stark, Greyjoy, Lannister, etc.)
type Faction struct {
	gorm.Model
	Name        string
	PowerTokens int
	Game        uint
	Cards       []Card `gorm:"foreignkey:Faction"`
	Areas       []Area `gorm:"foreignkey:Area"`
}
