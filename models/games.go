package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

//Game struct (FirstGame, SecondGame, etc...)
type Game struct {
	gorm.Model
	UserCreator uint
	Players     []User `gorm:"many2many:joined_games"`
	Name        string
	Active      bool
	Factions    []Faction
	Tracks      []Track
}

// AssignFactions creates a faction for each user in the game
func (game Game) AssignFactions(db *gorm.DB) error {

	if len(game.Players) < 6 {
		return fmt.Errorf("Not enough players to start this game")
	}

	factionNames := []string{"Stark", "Greyjoy", "Lannister", "Barratheon", "Martell", "Tyrell"}

	factionNames = shuffle(factionNames)

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

func shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
