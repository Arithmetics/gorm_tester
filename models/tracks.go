package models

import "github.com/jinzhu/gorm"

//Track is the power track (Sword, Throne, Raven)
type Track struct {
	gorm.Model
	Name        string
	Game        Game
	GameID      uint
	Bids        []Bid
	Position1   Faction
	Position1ID uint
	Position2   Faction
	Position2ID uint
	Position3   Faction
	Position3ID uint
	Position4   Faction
	Position4ID uint
	Position5   Faction
	Position5ID uint
	Position6   Faction
	Position6ID uint
}
