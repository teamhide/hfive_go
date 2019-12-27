package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserUsecase
}

func (uc UserController) RegisterDefaultUser(c *gin.Context) {
	var form RegisterDefaultUserRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uc.RegisterUserUsecase(form.Email, form.Password1, form.Password2)
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (uc UserController) KakaoLogin(c *gin.Context) {
	code := c.Query("code")
	token, refreshToken := KakaoLoginUsecase(code)
	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}

func (uc UserController) GoogleLogin(c *gin.Context) {
	code := c.Query("code")
	token, refreshToken := GoogleLoginUsecase(code)
	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}

func (uc UserController) RefreshToken(c *gin.Context) {
	var form RefreshTokenRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := RefreshTokenUsecase(form.Token, form.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc UserController) VerifyToken(c *gin.Context) {
	token := c.Param("token")
	if VerifyTokenUsecase(token) == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}
