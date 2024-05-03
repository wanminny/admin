package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Passwd   string `json:"passwd"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	isOn     bool   `json:"is_on"`
	Avatar   string `json:"avatar"`
}

//func (u *User) TableName() string {
//	return "zxw_user"
//}
