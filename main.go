package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teamhide/hfive_go/db"
	"github.com/teamhide/hfive_go/users"
)

func main() {
	r := gin.Default()
	db.Init()

	users := new(users.UserController)
	v1 := r.Group("/v1/users")
	{
		v1.GET("/", users.RegisterDefaultUser)
		v1.GET("/google/login", users.GoogleLogin)
		v1.GET("/kakao/login", users.KakaoLogin)
		v1.POST("/refresh", users.RefreshToken)
		v1.POST("/verify", users.VerifyToken)
	}
	r.Run(":8000")
}
