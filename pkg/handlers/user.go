package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wanminny/admin/pkg/global"
	"github.com/wanminny/admin/pkg/model"
	"github.com/wanminny/admin/pkg/util"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
)

type User struct {
	Db *gorm.DB
}

func NewUser() *User {
	return &User{
		Db: global.Db,
	}
}

type ReqUser struct {
	Name   string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Passwd string `form:"passwd" json:"passwd" binding:"required,min=3,max=10"`
}

func (u *User) GetUser(c *gin.Context) {
	r := util.NewResponse()
	id := c.Param("id")
	klog.Infoln(id, u.Db)
	user := model.User{}
	u.Db.First(&user, "id = ?", id)
	klog.Infoln(user)
	r.Data = user
	util.SetSuccess(c, r)
}

func (u *User) All(c *gin.Context) {
	r := util.NewResponse()
	users := []model.User{}
	//var users []model.User
	u.Db.Find(&users)
	r.Data = users
	util.SetSuccess(c, r)
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
