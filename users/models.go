package users

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID           uint64 `gorm:"primary_key; AUTO_INCREMENT"`
	Email        string `gorm:"size: 255"`
	Password     string `gorm:"size: 255"`
	LoginType    string `gorm:"size: 20"`
	AccessToken  string `gorm:"size: 255; unique"`
	RefreshToken string `gorm:"size: 255; unique"`
}
