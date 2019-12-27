package main

import (
	"github.com/gin-gonic/gin"
	users "github.com/teamhide/hfive_go/users/controllers"
)

func main() {
	r := gin.Default()
	// db.Init()

	users := new(users.UserController)
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", users.RegisterDefaultUser)
	}
	r.Run(":8000")
}
