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

func (game Game) CreateTracks (db *gorm.DB) error {
	track1 := Track{
		Name: "IronThrone",
		GameID: game.ID,
	}
	track2 := Track{
		Name: "Fiefdoms",
		GameID: game.ID,
	}
	track3 := Track{
		Name: "KingsCourt",
		GameID: game.ID
	}
	
	db.Save(&track1)
	db.Save(&track2)
	db.Save(&track3)
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
