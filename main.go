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
	// createUsersAndGame(db)
	// *******************************
	// allUsersJoinGame(db)
	// *******************************
	// setupGame(db)
	// *******************************
	// bidOnIronThrone(db)
	bids := []models.Bid{}
	db.Debug().Find(&bids)
	fmt.Println(len(bids))

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

func bidOnIronThrone(db *gorm.DB) {
	var factions []models.Faction
	db.Find(&factions)
	var game models.Game
	db.Preload("Tracks").First(&game)

	for i := 0; i < 6; i++ {
		fmt.Println(factions[i].MakeBid(i, "IronThrone", db))
	}

	var track models.Track

	db.Where(models.Track{Name: "IronThrone"}).First(&track)
	fmt.Printf("%+v\n", track.ID)
	fmt.Println(track.SettleBids(db))
}
