package controllers

import (
	"html/template"
	"./../models"
	identify "./../verify/auth"
)

type UserController struct {
	BaseController
}

var rules = map[string]int{
	"Login":   0,
	"Logout": identify.Login,
}

func (this *UserController) getRules(action string) int {
	return rules[action]
}

func (this *UserController) Login() {
	if (this.isUserLogin()) {
		//if has login,then go home
		this.Redirect("/", 302)
		return
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "user/index.tpl"
}

func (this *UserController) Logout() {
	this.LogoutUser()
	this.Redirect("/account/login", 302)
}

func (this *UserController) LoginPost() {
	if (this.isUserLogin()) {
		//if has login,then go home
		this.Redirect("/", 302)
		return
	}
	username := this.GetString("username")
	password := this.GetString("password")
	user := models.User{Name:username, Password: password}
	if user.LoginVerify() {
		//验证通过
		this.LoginUser(user.ID, user.Name);
		this.Redirect("/", 302)
		return
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["username"] = username
	this.TplName = "user/index.tpl"
	//this.TplName = "user/index.tpl"
}