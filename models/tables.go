package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey autoIncrement"`
	Token        string `gorm:"unique"`
	RefreshToken string `gorm:"unique"`
	TimeToLive   int    `gorm:"not null"`
	UserId	 	 string `gorm:"unique not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

