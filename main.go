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
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Game{})

	// Create Users
	// db.Create(&models.User{Name: "Bob", Rank: "Basic"})
	// db.Create(&models.User{Name: "Dave", Rank: "Basic"})
	// db.Create(&models.User{Name: "Steve", Rank: "Basic"})
	// db.Create(&models.User{Name: "Brock", Rank: "Advanced"})

	var userFirst models.User
	var userLast models.User
	db.Preload("JoinedGames.Players").First(&userFirst)
	db.Last(&userLast)

	// Create Items
	// db.Create(&models.Item{Name: "Shield", Price: 34, Action: "Defend", UserBuyer: userFirst.ID, UserSeller: userLast.ID})
	// db.Create(&models.Item{Name: "Sword", Price: 22, Action: "Attack", UserBuyer: userLast.ID, UserSeller: userFirst.ID})

	//Create Games
	// newGame := &models.Game{Name: "FunGame", Active: true, UserCreator: userFirst.ID, Players: []models.User{userFirst, userLast}}

	// db.Create(&newGame)
	fmt.Println(userFirst.JoinedGames[0].Players)
}
