package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Surname string
	Name    string
	Email   string
	Token   string
}

type Session struct {
	gorm.Model
	UserID    uint `gorm:"foreignKey:UserID;references:ID"`
	Token     string
	ExpiresAt time.Time
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
	UserID        uint `gorm:"foreignKey:UserID;references:ID"`
	Name          string
	Price         uint
	Quantity      uint
	Token         string
}

type Order struct {
	gorm.Model
	CompositionID string
	UserID        uint
}
