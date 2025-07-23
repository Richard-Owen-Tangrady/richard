package models

type Cart struct {
	CartID    string `gorm:"unique" json:"cart_id"`
	UserID    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
