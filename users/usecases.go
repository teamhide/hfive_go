package users

import (
	"errors"
	"github.com/teamhide/gin_boilerplate/db"
	"github.com/teamhide/gin_boilerplate/models"
)

type UserUsecase interface {
	RegisterUserUsecase(email, password1, password2 string) (bool, error)
	GoogleLoginUsecase(code string) (string, string)
	KakaoLoginUsecase(code string) (string, string)
	RefreshTokenUsecase(token string, refreshToken string) string
	VerifyTokenUsecase(token string) bool
}

func RegisterUserUsecase(email, password1, password2 string) (bool, error) {
	db := db.GetDB()
	var user models.User

	if password1 != password2 {
		return false, errors.New("password1 and password2 does not match")
	}
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return false, errors.New("get user errors")
	}
	return true, nil
}

func GoogleLoginUsecase(code string) (string, string) {
	return "a", "b"
}

func KakaoLoginUsecase(code string) (string, string) {
	return "a", "b"
}

func RefreshTokenUsecase(token string, refreshToken string) string {
	return "a"
}

func VerifyTokenUsecase(token string) bool {
	return true
}
