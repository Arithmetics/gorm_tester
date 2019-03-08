package models

import "github.com/jinzhu/gorm"

//Area struct (The Reach, Winterfell, etc..)
type Area struct {
	gorm.Model
	Name    string
	Used    bool
	Faction uint
}
