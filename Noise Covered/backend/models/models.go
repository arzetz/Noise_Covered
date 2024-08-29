package models

import (
	"time"

	"gorm.io/gorm"
)

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
	CompositionID uint `gorm:"foreignKey:CompositionID;references:ID"`
	CreatedAt     time.Time
	ExpiresAt     time.Time
	Name          string
	Price         uint
	Quantity      uint
}

type Order struct {
	gorm.Model
	CompositionID string
	UserID        uint
}
