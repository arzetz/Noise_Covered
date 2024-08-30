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

type Session struct {
	gorm.Model
	UserID    string
	Token     string
	ExpiresAT time.Time
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
	UserID        string
	Name          string
	Price         uint
	Quantity      uint
}

type Order struct {
	gorm.Model
	CompositionID string
	UserID        uint
}
