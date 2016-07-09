package controllers

import (
	"github.com/astaxie/beego"
	identify "./../verify/auth"
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
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {
	var _, action = this.GetControllerAndAction()
	if app, ok := this.AppController.(Rules); ok {
		rule := app.getRules(action)
		if ((rule & identify.Login) == identify.Login) && !this.IsUserLogin() {
			if rule & identify.JumpBack == identify.JumpBack {
				//"&query=" + this.Ctx.Request.URL.RawQuery
				this.Redirect("/account/signin?next=" + this.Ctx.Request.URL.Path, 302)
			} else {
				this.Redirect("/account/signin", 302)
			}
		}
	}
}

func (this *BaseController)IsUserLogin() bool {
	s := this.GetSession(Is_Login)
	if s == nil {
		return false
	}
	//fmt.Println(s)
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

func (this *BaseController) getUserId() uint {
	return (this.GetSession(User)).(uint)
}
func (this *BaseController) getUsername() string {
	return (this.GetSession(User_Name)).(string)
}