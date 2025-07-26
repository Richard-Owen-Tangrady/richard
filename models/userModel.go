package models

import "time"

type User struct {
	UserID      string    `gorm:"primaryKey;type:varchar(255)" json:"user_id"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type Body struct {
	Email    string `gorm:"unique" json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
