package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) RegisterDefaultUser(c *gin.Context) {
	var form RegisterDefaultUserRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	RegisterUserUsecase(form.Email, form.Password1, form.Password2)
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (u UserController) KakaoLogin(c *gin.Context) {
	code := c.Query("code")
	token, refreshToken := KakaoLoginUsecase(code)
	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}

func (u UserController) GoogleLogin(c *gin.Context) {
	code := c.Query("code")
	token, refreshToken := GoogleLoginUsecase(code)
	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}

func (u UserController) RefreshToken(c *gin.Context) {
	var form RefreshTokenRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := RefreshTokenUsecase(form.Token, form.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u UserController) VerifyToken(c *gin.Context) {
	token := c.Param("token")
	if VerifyTokenUsecase(token) == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}
