package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) RegisterDefaultUser(c *gin.Context) {
	RegisterUserUsecase()
	c.JSON(http.StatusCreated, gin.H{"user": "hello"})
}

func (u UserController) KakaoLogin(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"kakaologin": "1"})
}

func (u UserController) GoogleLogin(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"Googlelogin": 1})
}

func (u UserController) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"refreshotken": "!"})
}

func (u UserController) VerifyToken(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"verifytoken": 1})
}
