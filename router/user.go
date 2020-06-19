package router

import (
	"demo/db"
	"demo/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func RegisterUser(c *gin.Context) {
	var (
		user model.User
		err  error
	)
	err = c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("user---", user)
	err = db.Engine.Create(&user).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册用户失败",
			"data": err.Error(),
		})
		return
	}
	data, _ := json.Marshal(user)
	db.GetRedis().Set("user_"+strconv.Itoa(user.Id), string(data), 0)
	c.JSON(http.StatusOK, gin.H{
		"msg": "用户注册成功",
	})
}

func GetUser(c *gin.Context) {
	name := c.PostForm("name")
	log.Println("name----", name)
	var user model.User
	db.Engine.Table("user").Where("name = ?", name).Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询用户成功",
		"data": user,
	})
}

func DelUser(c *gin.Context) {
	name := c.PostForm("name")
	var user model.User
	db.Engine.Table("user").Where("name = ?", name).Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除用户成功",
	})
}
