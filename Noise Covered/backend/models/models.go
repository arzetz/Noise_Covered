package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Surname string
	Name    string
	Email   string `gorm:"unique"`
}

type Composition struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Author string
	Album  string
	Genre  string
	Price  uint
}

type Basket struct {
	gorm.Model
	CompositionID uint `gorm:"foreignKey:CompositionID;references:ID"`
	Name          string
	Price         uint
	Quantity      uint
}

type Order struct {
	gorm.Model
	CompositionID string
	UserID        uint
}
