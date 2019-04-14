package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//User struct (Arithmetics, MrCream, etc...)
type User struct {
	gorm.Model
	Name         string
	Rank         string
	CreatedGames []Game `gorm:"foreignkey:UserCreator"`
	JoinedGames  []Game `gorm:"many2many:joined_games"`
}

// CreateGame makes a new game with the User as the creator
func (user User) CreateGame(gameName string, db *gorm.DB) (uint, error) {
	newGame := Game{
		Name:        gameName,
		UserCreator: user.ID,
		Active:      true,
		Players:     []User{user},
	}
	err := db.Create(&newGame).Error

	newGame.CreateTracks(db)

	return newGame.ID, err
}

// JoinGame is used to join an existing game as a player
func (user User) JoinGame(gameID uint, db *gorm.DB) error {
	var game Game
	db.Preload("Players").First(&game, gameID)

	if len(game.Players) > 5 {
		return fmt.Errorf("No open spots in this game")
	}

	// need to check if player is already in game here as well
	for _, player := range game.Players {
		if player.ID == user.ID {
			return fmt.Errorf("User has already joined the game")
		}
	}

	game.Players = append(game.Players, user)

	err := db.Save(&game).Error

	return err
}

//Faction For a certain user gets the Faction they are playing in that game
func (user User) Faction(gameID uint, db *gorm.DB) *Faction {
	var faction Faction

	db.Preload("Bid").Preload("Game.Factions").Where(&Faction{UserID: user.ID, GameID: gameID}).First(&faction)

	return &faction
}
