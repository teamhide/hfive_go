package usecases

import (
	"errors"

	"github.com/teamhide/hfive_go/db"
	"github.com/teamhide/hfive_go/users/models"
)

type UserUsecase struct {
}

func (u UserUsecase) RegisterUserUsecase(email, password1, password2 string) (bool, error) {
	db := db.GetDB()
	userModel := models.User{}

	if password1 != password2 {
		return false, errors.New("password1 and password2 does not match")
	}
	if !db.Where("email = ?", email).First(&userModel).RecordNotFound() {
		return false, errors.New("user already exists")
	}

	userModel.Email = email
	userModel.Password = password1
	userModel.LoginType = "Default"
	if err := db.Create(&userModel); err != nil {
		return false, errors.New("create user errors")
	}
	return true, nil
}

func (u UserUsecase) GoogleLoginUsecase() (bool, error) {
	return true, nil
}

func (u UserUsecase) KakaoLoginUsecase() (bool, error) {
	return true, nil
}

func (u UserUsecase) RefreshTokenUsecase() (bool, error) {
	return true, nil
}

func (u UserUsecase) VerifyTokenUsecase() (bool, error) {
	return true, nil
}
