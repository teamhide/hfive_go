package users

import "time"

type User struct {
	ID           uint   `gorm:"primary_key"`
	Email        string `gorm:"size: 255"`
	Password     string `gorm:"size: 255"`
	LoginType    string `gorm:"size: 20"`
	AccessToken  string `gorm:"size: 255; unique"`
	RefreshToken string `gorm:"size: 255; unique"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
