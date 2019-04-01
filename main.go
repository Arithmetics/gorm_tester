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
	/////////////  Create Users
	// db.Create(&models.User{Name: "Bob", Rank: "Basic"})
	// db.Create(&models.User{Name: "Dave", Rank: "Basic"})
	// db.Create(&models.User{Name: "Sara", Rank: "Basic"})
	// db.Create(&models.User{Name: "Monkee", Rank: "Advanced"})

	// var user models.User
	// db.First(&user)

	// fmt.Println(user.CreateGame("CoolGame", db))

	// *******************************
	////////////  Join Game

	// var user models.User

	// db.Find(&user, 3)
	// user.JoinGame(1, db)
	// db.Find(&user, 4)
	// user.JoinGame(1, db)
	// db.Find(&user, 5)
	// user.JoinGame(1, db)
	// db.Find(&user, 6)
	// user.JoinGame(1, db)

	// *****
	//////// SEtup game
	var game models.Game
	db.Preload("Players").First(&game)
	game.AssignFactions(db)
}
