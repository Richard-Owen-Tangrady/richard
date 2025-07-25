package models

import "time"

type Cart struct {
	CartID       string  `gorm:"primaryKey" json:"cart_id"`
	UserRefer    string  `gorm:"type:varchar(255)" json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	ProductRefer string  `gorm:"type:varchar(255)" json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	Quantity     int     `json:"quantity"`
	CreatedAt    time.Time
}

type CreteCart struct {
	UserRefer    string `json:"user_id"`
	ProductRefer string `json:"product_id"`
	Quantity     int    `json:"quantity"`
}
