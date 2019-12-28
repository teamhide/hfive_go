package usecases

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/teamhide/hfive_go/core/configs"
	"github.com/teamhide/hfive_go/core/db"
	"github.com/teamhide/hfive_go/users/models"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
}

func (u UserUsecase) RegisterUserUsecase(email, password1, password2 string) (string, error) {
	db := db.GetDB()
	config := configs.GetConfiguration()
	userModel := models.User{}

	if password1 != password2 {
		return "", errors.New("password1 and password2 does not match")
	}
	if !db.Where("email = ?", email).First(&userModel).RecordNotFound() {
		return "", errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.MinCost)
	if err != nil {
		return "", errors.New("hashing password error")
	}

	userModel.Email = email
	userModel.Password = string(hash)
	userModel.LoginType = "Default"
	db.Create(&userModel)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userModel.ID,
		"exp":     time.Now().Add(time.Minute * 60),
	})
	tokenString, err := token.SignedString([]byte(config.JwtSecretKey))
	return tokenString, nil
}

func (u UserUsecase) GoogleLoginUsecase(code string) (string, error) {
	return "", nil
}

func (u UserUsecase) KakaoLoginUsecase(code string) (string, error) {
	return "", nil
}

func (u UserUsecase) RefreshTokenUsecase(token, refreshToken string) (string, error) {
	return "", nil
}

func (u UserUsecase) VerifyTokenUsecase(token string) (bool, error) {
	return true, nil
}
