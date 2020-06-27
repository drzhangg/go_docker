package main

import (
	"demo/config"
	"demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api")
	{
		v1.POST("/register", router.RegisterUser)
		v1.POST("/getUser", router.GetUser)
		v1.POST("/deleteUser", router.DelUser)
	}
	config.DBConfig()
	r.Run(":8080")
}
