package models

import "github.com/jinzhu/gorm"

//Game struct (FirstGame, SecondGame, etc...)
type Game struct {
	gorm.Model
	UserCreator uint
	Players     []User `gorm:"many2many:joined_games"`
	Name        string
	Active      bool
	Factions    []Faction
}

// AssignFactions creates a faction for each user in the game
func (game Game) AssignFactions(db *gorm.DB) error {
	factionNames := []string{"Stark", "Greyjoy", "Lannister", "Barratheon", "Martell", "Tyrell"}

	for i, name := range factionNames {
		faction := Faction{
			Name:        name,
			PowerTokens: 10,
			GameID:      game.ID,
			User:        game.Players[i],
		}

		db.Create(&faction)
	}

	return nil
}
