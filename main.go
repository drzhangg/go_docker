package main

import (
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
	r.Run(":8080")
}
