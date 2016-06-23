package models

import (
	"github.com/astaxie/beego/validation"
	"io"
	"crypto/sha1"
	//"encoding/hex"
	"github.com/jinzhu/gorm"
	"./database"
	//"encoding/hex"
	"encoding/hex"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func (User) TableName() string {
	return "user"
}

func (u *User) validPassword(v *validation.Validation) {
	h := sha1.New()
	io.WriteString(h, u.Password)
	s := hex.EncodeToString(h.Sum(nil))
	database.DB.Where("tel = ? AND password_hash = ?", u.Name, s).First(u)
	if u.ID == 0 {
		v.SetError("name", "用户名或密码错误")
		v.SetError("pass", "用户名或密码错误")
	}
}

func (this *User)LoginVerify() ([]*validation.Error,bool) {
	valid := validation.Validation{}
	valid.Required(this.Name, "name").Message("用户名不能为空")
	valid.Required(this.Password, "pass").Message("密码不能为空")
	if valid.HasErrors() {
		return valid.Errors,false
	}

	this.validPassword(&valid); //验证密码
	if valid.HasErrors() {
		return valid.Errors,false
	}
	return nil,true
}

func (this *User) Logout() bool {
	return false
}
