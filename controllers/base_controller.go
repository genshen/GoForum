package controllers

import (
	"github.com/astaxie/beego"
	identify "beeapp/verify/auth"
	"fmt"
)

const (
	User = "user"
	User_Name = "username"
	Is_Login = "is_login"
)

type Rules interface {
	getRules(string) int
}

type BaseController struct {
	beego.Controller
	isLogin bool
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {
	var _, action = this.GetControllerAndAction()
	if app, ok := this.AppController.(Rules); ok {
		if (app.getRules(action) & identify.Login) != 0 && !this.isUserLogin() {
			//not login
			this.Redirect("/account/login", 302)
		}
	}
}

func (this *BaseController)isUserLogin() bool {
	s := this.GetSession(Is_Login)
	if s == nil {
		return false
	}
	fmt.Println(s)
	return s.(bool)
}

func (this *BaseController) LoginUser(ID uint, Name string) bool {
	this.SetSession(Is_Login, true)
	this.SetSession(User, ID)
	this.SetSession(User_Name, Name)
	return true
}

func (this *BaseController) LogoutUser() bool {
	this.DelSession(Is_Login)
	this.DelSession(User)
	this.DelSession(User_Name)
	return true
}