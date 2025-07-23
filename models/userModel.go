package models

import "time"

type User struct {
	UserID      string    `gorm:"unique" json:"user_id"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type Body struct {
	UserID   string `gorm:"unique" json:"user_id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
