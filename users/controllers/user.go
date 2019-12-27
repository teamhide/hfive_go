package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamhide/hfive_go/users/serializers"
	"github.com/teamhide/hfive_go/users/usecases"
)

type UserController struct {
	usecases.UserUsecase
}

func (u UserController) RegisterDefaultUser(c *gin.Context) {
	var form serializers.RegisterDefaultUserRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := u.RegisterUserUsecase(form.Email, form.Password1, form.Password2)
	if status == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}
