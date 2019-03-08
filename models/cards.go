package models

import "github.com/jinzhu/gorm"

//Card struct (Ramsay, The Mountain, etc..)
type Card struct {
	gorm.Model
	Name    string
	Used    bool
	Faction uint
}
