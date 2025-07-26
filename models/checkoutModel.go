package models

import "time"

type Checkout struct {
	CheckoutID   string    `gorm:"primaryKey" json:"checkout_id"`
	ProductRefer string    `gorm:"type:varchar(255)" json:"product_id"`
	Product      Product   `gorm:"foreignKey:ProductRefer"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
}

type CheckoutCreateRequest struct {
	ProductRefer string `gorm:"type:varchar(255)" json:"product_id"`
	Quantity     int    `json:"quantity"`
}
