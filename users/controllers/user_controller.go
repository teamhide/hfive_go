package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	users "github.com/teamhide/hfive_go/users/usecases"
)

type UserController struct {
	users.UserUsecase
}

func (u UserController) RegisterDefaultUser(c *gin.Context) {
	u.RegisterUserUsecase("padocon@naver.com", "123", "123")
	c.JSON(http.StatusOK, gin.H{"status": true})
}
