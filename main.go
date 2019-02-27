package main

import (
	"fmt"

	"github.com/arithmetics/gorm_tester/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Product struct
type Product struct {
	gorm.Model
	Code    string
	Price   uint
	StoreID uint
}

//Store struct
type Store struct {
	gorm.Model
	City  string
	Level uint
}

func main() {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "brocktillotson", "gorm", "uiop789&*()")

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Create
	db.Create(&models.User{Name: "Bob", Rank: "Basic"})

}
