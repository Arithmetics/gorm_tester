package main

import (
	"fmt"

	"github.com/arithmetics/gorm_tester/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "brocktillotson", "gorm", "uiop789&*()")

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Bid{})
	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.Faction{})
	db.AutoMigrate(&models.Track{})

	// *******************************
	///////////  Create Users
	// db.Create(&models.User{Name: "Bob", Rank: "Basic"})
	// db.Create(&models.User{Name: "Dave", Rank: "Basic"})
	// db.Create(&models.User{Name: "Sara", Rank: "Basic"})
	// db.Create(&models.User{Name: "Monkee", Rank: "Advanced"})
	// db.Create(&models.User{Name: "Feez", Rank: "Advanced"})
	// db.Create(&models.User{Name: "Donkey", Rank: "Advanced"})

	// var user models.User
	// db.First(&user)

	// fmt.Println(user.CreateGame("CoolGame", db))

	// *******************************
	////////////  Join Game
	var game models.Game
	var user models.User

	db.First(&game)

	db.Find(&user, 7)
	user.JoinGame(game.ID, db)
	db.Find(&user, 8)
	user.JoinGame(game.ID, db)
	db.Find(&user, 9)
	user.JoinGame(game.ID, db)
	db.Find(&user, 10)
	user.JoinGame(game.ID, db)
	db.Find(&user, 11)
	user.JoinGame(game.ID, db)
	db.Find(&user, 12)
	user.JoinGame(game.ID, db)
	// *****
	//////// Setup game
	// db.Preload("Players").First(&game)
	// game.AssignFactions(db)

	/// ****************
	////// Get faction for the player in the game

	// var game models.Game
	// var user models.User

	// db.First(&user)
	// db.First(&game)
	// fmt.Println(user.ID)
	// fmt.Println(game.ID)

	// faction := user.Faction(game.ID, db)

	// faction.MakeBid(2)
}
