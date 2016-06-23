package controllers

import (
	"html/template"
	"./../models"
	identify "./../verify/auth"
	"../verify/form"
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

func (this *UserController) SignIn() {
	if (this.isUserLogin()) {
		//if has login,then go home
		this.Redirect("/", 302)
		return
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signin.html"
}

func (this *UserController) SignInPost() {
	if (this.isUserLogin()) {
		//if has login,then go home
		this.Redirect("/", 302)
		return
	}
	username := this.GetString("username")
	password := this.GetString("password")
	user := models.User{Name:username, Password: password}
	if errs, status := user.LoginVerify(); status {
		//验证通过
		this.LoginUser(user.ID, user.Name)
		this.Redirect("/", 302)
		return
	} else {
		this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
		s := form.NewInstant(errs, map[string]string{"name":  username, "pass": ""})
		this.Data["form_check"] = string(s)
		this.TplName = "account/signin.html"
	}
}

func (this *UserController) SignUp(){
	if (this.isUserLogin()) {
		this.Redirect("/", 302)
		return
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signup.html"
}

func (this *UserController) SignUpPost(){
	if (this.isUserLogin()) {
		this.Redirect("/", 302)
		return
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signup.html"
}

func (this *UserController) SignOut() {
	this.LogoutUser()
	this.Redirect("/account/signin", 302)
}
