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

	// create tracks here, maybe move to method on game??
	track1 := Track{
		Name:   "IronThrone",
		GameID: newGame.ID,
	}

	track2 := Track{
		Name:   "Fiefdoms",
		GameID: newGame.ID,
	}

	track3 := Track{
		Name:   "KingsCourt",
		GameID: newGame.ID,
	}

	db.Create(&track1)
	db.Create(&track2)
	db.Create(&track3)

	if err != nil {
		return 0, err
	}

	return newGame.ID, nil
}

// JoinGame is used to join an existing game as a player
func (user User) JoinGame(gameID uint, db *gorm.DB) error {
	var game Game
	db.First(&game, gameID)

	if len(game.Players) > 5 {
		return fmt.Errorf("No open spots in this game")
	}

	// need to check if player is already in game here as well

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
