package controllers

import (
	"html/template"
	form_check "../verify/form"
	"./../models/forms"
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

func (this *UserController) SignIn() {
	if (this.isUserLogin()) {
		//if has login,then go home
		this.Redirect("/", 302)
		return
	}
	if this.Ctx.Request.Method == "POST" {
		username := this.GetString("username")
		password := this.GetString("password")
		sign_in_form := forms.SignInForm{Username:username, Password: password}
		if errs, status, userID := sign_in_form.LoginVerify(); status {
			//验证通过
			this.LoginUser(userID, username)
			this.Redirect("/", 302)
			return
		} else {
			s := form_check.NewInstant(errs, map[string]string{"email":  username, "pass": ""})
			this.Data["form_check"] = string(s)
		}
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signin.html"
}

func (this *UserController) SignUp() {
	if (this.isUserLogin()) {
		this.Redirect("/", 302)
		return
	}
	if this.Ctx.Request.Method == "POST" {
		email := this.GetString("email")
		password := this.GetString("password")
		sign_up_form := forms.SignUpForm{Email:email, Password: password}
		if errs, status := sign_up_form.Valid(); status {
			this.Data["email"] = email
			this.TplName = "account/signup_success.html"
			return
		} else {
			s := form_check.NewInstant(errs, map[string]string{"email":  email, "password": ""})
			this.Data["form_check"] = string(s)
		}
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signup.html"
}

func (this *UserController) SignOut() {
	this.LogoutUser()
	this.Redirect("/account/signin", 302)
}
