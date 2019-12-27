package routers

import "github.com/teamhide/hfive_go/users/controllers"

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	users := new(controllers.UserController)
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", users.RegisterDefaultUser)
	}
}
