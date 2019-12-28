package routers

import "github.com/teamhide/hfive_go/users/controllers"

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	users := new(controllers.UserController)
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", users.RegisterDefaultUser)
		r.GET("/oauth/google/login", users.GoogleLogin)
		r.GET("/oauth/kakao/login", users.KakaoLogin)
		v1.POST("/refresh", users.RefreshToken)
		v1.POST("/verify", users.VerifyToken)
	}
}
