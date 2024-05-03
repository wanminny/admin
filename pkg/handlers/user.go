package handlers

import (
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUser(c *gin.Context) {
	id := c.Param("id")
	userId := c.Query("userId")
	age := c.DefaultQuery("age", "0") // 如果找不到age参数，则使用默认值0
	c.JSON(200, gin.H{
		"message": "pong",
		"id":      id,
		"userId":  userId,
		"age":     age,
	})
}

func (u *User) Reg(c *gin.Context) {
	id := c.Param("id")
	userId := c.Query("userId")
	age := c.DefaultQuery("age", "0") // 如果找不到age参数，则使用默认值0
	c.JSON(200, gin.H{
		"message": "pong",
		"id":      id,
		"userId":  userId,
		"age":     age,
	})
}
