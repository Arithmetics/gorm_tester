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
	createUsersAndGame(db)
	// *******************************
	allUsersJoinGame(db)
	// *******************************

	//////// Setup game
	// db.Preload("Players").First(&game)
	// game.AssignFactions(db)

	/// ****************
	////// Get faction for the player in the game

	// var game models.Game
	// var user models.User

	// db.Find(&user, 12)
	// db.First(&game)
	// fmt.Println(user.ID)
	// fmt.Println(game.ID)

	// faction := user.Faction(game.ID, db)

	// fmt.Println(faction.MakeBid(5, "IronThrone", db))

	// var bid models.Bid
	// db.First(&bid)

	// fmt.Printf("%+v", bid)

	var track models.Track
	db.First(&track)

	// fmt.Println(track.SettleBids(db))

	fmt.Printf("%+v", &track)

}

func createUsersAndGame(db *gorm.DB) {
	db.Create(&models.User{Name: "Bob", Rank: "Basic"})
	db.Create(&models.User{Name: "Dave", Rank: "Basic"})
	db.Create(&models.User{Name: "Sara", Rank: "Basic"})
	db.Create(&models.User{Name: "Monkee", Rank: "Advanced"})
	db.Create(&models.User{Name: "Feez", Rank: "Advanced"})
	db.Create(&models.User{Name: "Donkey", Rank: "Advanced"})

	var user models.User
	db.First(&user)

	fmt.Println(user.CreateGame("CoolGame", db))
}

func allUsersJoinGame(db *gorm.DB) {
	var game models.Game
	db.First(&game)

	users := []models.User{}
	db.Find(&users)

	for _, user := range users {
		user.JoinGame(game.ID, db)
	}
}

func setupGame(db *gorm.DB) {
	var game models.Game

	db.Preload("Players").First(&game)

	game.AssignFactions(db)
}
