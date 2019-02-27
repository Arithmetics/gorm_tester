package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Item struct: belongs to User
type Item struct {
	gorm.Model
	UserSeller uint
	UserBuyer  uint
	Name       string
	Price      uint
	Action     string
}

//PrettyPrint prints full item to console
func (item Item) PrettyPrint() {
	fmt.Println("---\nItem:")
	fmt.Printf("Seller: %v\n", item.UserSeller)
	fmt.Printf("Buyer: %v\n", item.UserBuyer)
	fmt.Printf("Name: %v\n", item.Name)
	fmt.Printf("Price: %v\n", item.Price)
	fmt.Printf("Action: %v\n", item.Action)
	fmt.Printf("\n")
}
