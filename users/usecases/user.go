package usecases

import (
	"fmt"

	"github.com/teamhide/hfive_go/db"
	"github.com/teamhide/hfive_go/users/models"
)

type UserUsecase struct {
}

func (u UserUsecase) RegisterUserUsecase(email, password1, password2 string) (bool, error) {
	db := db.GetDB()
	fmt.Println(db.First(&models.User{}))
	return true, nil
}
