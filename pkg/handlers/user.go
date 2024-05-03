package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wanminny/admin/pkg/global"
	"github.com/wanminny/admin/pkg/model"
	"github.com/wanminny/admin/pkg/util"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

type ReqUser struct {
	Name   string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Passwd string `form:"passwd" json:"passwd" binding:"required,min=3,max=10"`
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
	r := util.NewResponse()
	var req ReqUser
	err := c.BindJSON(&req)
	if err != nil {
		util.SetFailed(c, r, err)
		return
	}
	db := global.Db
	user := model.User{
		Name:   req.Name,
		Passwd: req.Passwd,
	}
	db.Create(&user)
	r.Data = req
	util.SetSuccess(c, r)
}

func (u *User) Modify(c *gin.Context) {
	r := util.NewResponse()
	var req ReqUser
	err := c.BindJSON(&req)
	if err != nil {
		util.SetFailed(c, r, err)
		return
	}
	r.Data = req
	util.SetSuccess(c, r)
}

// Delete http://localhost:8090/v1/user/1?name=abc&passwd=10
func (u *User) Delete(c *gin.Context) {
	r := util.NewResponse()
	var req ReqUser
	err := c.ShouldBindQuery(&req)
	if err != nil {
		util.SetFailed(c, r, err)
		return
	}
	r.Data = req
	util.SetSuccess(c, r)
}
