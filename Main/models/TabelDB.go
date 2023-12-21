package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName string `gorm:"not null; not empty" json:"full_name" validate:"required"`
	Email    string `gorm:"unique; not null; not empty" json:"email" validate:"required,email"`
	Password string `gorm:"not null; not empty" json:"password" validate:"required,min=5"`
	Role     string `gorm:"not null; not empty; default:'customer'" json:"role"`
	Balance  int    `gorm:"not null; not empty; default:0" json:"balance" validate:"gte=0,lte=100000000"`
}

type Products struct {
	gorm.Model
	Title      string `gorm:"not null; not empty" json:"title" validate:"required"`
	Price      int    `gorm:"not null; not empty" json:"price" validate:"gte=5"`
	Stock      int    `gorm:"not null; not empty" json:"stock" validate:"gte=0,lte=50000000"`
	CategoryID int    `gorm:"foreignKey:CategoryID;references:ID" json:"category_id"`
}

type Categories struct {
	gorm.Model
	Type        string     `gorm:"not null; not empty" json:"type" validate:"required"`
	SoldProduct int        `gorm:"default:0" json:"sold_product_amount"`
	Products    []Products `gorm:"foreignKey:CategoryID;references:ID" json:"Products"`
}

type TransactionHistories struct {
	gorm.Model
	ProductID  int      `gorm:"foreignKey:ProductID" json:"product_id"`
	UserID     int      `gorm:"foreignKey:UserID" json:"user_id"`
	Quantity   int      `gorm:"not null; not empty" json:"quantity" validate:"required"`
	TotalPrice int      `gorm:"not null; not empty" json:"total_price" validate:"required"`
	Product    Products `gorm:"foreignKey:ProductID;references:ID" json:"Product"`
	User       Users    `gorm:"foreignKey:UserID;references:ID" json:"User"`
}
