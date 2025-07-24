package models

import "time"

type Cart struct {
	CartID       string  `gorm:"primaryKey" json:"cart_id"`
	UserRefer    string  `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	ProductRefer string  `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	Quantity     int     `json:"quantity"`
	CreatedAt    time.Time
}

type CreteCart struct {
	CartID   string  `gorm:"unique" json:"cart_id"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}
