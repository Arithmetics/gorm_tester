package models

import "github.com/jinzhu/gorm"

//User struct
type User struct {
	gorm.Model
	Name        string
	Rank        string
	SoldItems   []Item `gorm:"foreignkey:UserSeller"`
	BoughtItems []Item `gorm:"foreignkey:UserBuyer"`
}
