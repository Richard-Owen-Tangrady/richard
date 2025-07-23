package models

type Product struct {
	ProductID   string `gorm:"unique" json:"product_id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type ProductCreateRequest struct {
	ProductID string `gorm:"unique" json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}
