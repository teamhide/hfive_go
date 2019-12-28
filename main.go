package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teamhide/hfive_go/core/db"
	"github.com/teamhide/hfive_go/routers"
)

func main() {
	r := gin.Default()
	db.Init()
	routers.Init(r)
	r.Run(":8000")
}
