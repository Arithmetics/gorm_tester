package models

import "github.com/jinzhu/gorm"

//Faction (Stark, Greyjoy, Lannister, etc.)
type Faction struct {
	gorm.Model
	Name        string
	PowerTokens int
	GameID      uint
	User        User
	UserID      uint
}
