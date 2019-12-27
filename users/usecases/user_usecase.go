package users

import (
	"github.com/teamhide/hfive_go/db"
	users "github.com/teamhide/hfive_go/users/models"
)

type UserUsecase struct {
}

func (u UserUsecase) RegisterUserUsecase(email, password1, password2 string) (bool, error) {
	db := db.GetDB()
	db.First(&users.User{})
	return true, nil
}
