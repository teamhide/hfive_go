package users

import (
	"github.com/teamhide/gin_boilerplate/db"
	"github.com/teamhide/gin_boilerplate/models"
)

func RegisterUserUsecase(email, password1, password2 string) (bool, string) {
	db := db.GetDB()
	var user models.User

	if password1 != password2 {
		return false, "password1 and password2 does not match"
	}
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return false, "get user error"
	}
	return true, ""
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
