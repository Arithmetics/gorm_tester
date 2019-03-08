package models

import "github.com/jinzhu/gorm"

//User struct (Arithmetics, MrCream, etc...)
type User struct {
	gorm.Model
	Name         string
	Rank         string
	SoldItems    []Item `gorm:"foreignkey:UserSeller"`
	BoughtItems  []Item `gorm:"foreignkey:UserBuyer"`
	CreatedGames []Game `gorm:"foreignkey:UserCreator"`
	JoinedGames  []Game `gorm:"many2many:joined_games"`
}

//Deficit calcs users profit
func (user User) Deficit() int {
	soldTotal := 0
	boughtTotal := 0

	for _, item := range user.SoldItems {
		soldTotal += int(item.Price)
	}

	for _, item := range user.BoughtItems {
		boughtTotal += int(item.Price)
	}

	return soldTotal - boughtTotal
}
