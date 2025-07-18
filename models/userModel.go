package models

import "time"

type User struct {
	UserID      string    `gorm:"unique" json:"userid"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phonenumber"`
	CreatedAt   time.Time `json:"createdat"`
}
