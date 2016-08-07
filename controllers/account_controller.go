package controllers

import (
	"html/template"
	"github.com/astaxie/beego"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/components/event"
	"gensh.me/goforum/components/context/account"
	identify "gensh.me/goforum/models/values"
)

type AccountController struct {
	BaseController
}

var rules = map[string]int{
	"Login":   0,
	"Logout": identify.Login,
}

type SignResult struct {
	Next   string `json:"next"`
	Status bool   `json:"status"`
}

func (this *AccountController) getRules(action string) int {
	return rules[action]
}

func (this *AccountController) SignIn() {
	if (this.IsUserLogin()) {
		//if has login,then go home
		this.Redirect("/", 302)
		return
	}
	//this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signin.html"
}

func (this *AccountController) POST_SignIn() {
	if (this.IsUserLogin()) {
		this.Redirect("/", 302)
		return
	}
	username := this.GetString("username")
	password := this.GetString("password")
	sign_in_form := account.SignInForm{Username:username, Password: password}
	if errs, userID := sign_in_form.SignInVerify(); errs == nil {
		//验证通过
		this.LoginUser(userID, sign_in_form.Username)
		next := this.GetString("next")
		if len(next) > 0 && next[0] != '/' {
			next = "/" + next
		}
		this.Data["json"] = &utils.SimpleJsonResponse{Status:1, Addition:next}
	} else {
		s := utils.NewInstant(errs, map[string]string{"name":  username, "pass": ""})
		this.Data["json"] = &utils.SimpleJsonResponse{Status:0, Error:&s}
	}
	this.ServeJSON()
}

func (this *AccountController) SignUp() {
	if (this.IsUserLogin()) {
		this.Redirect("/", 302)
		return
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "account/signup.html"
}

func (this *AccountController) POST_SignUp() {
	if (this.IsUserLogin()) {
		this.Redirect("/", 302)
		return
	}
	email := this.GetString("email")
	nickname := this.GetString("nickname")
	password := this.GetString("password")
	sign_up_form := account.SignUpForm{Email:email, Nickname:nickname, Password: password}
	if errs := sign_up_form.Valid(); errs == nil {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:1}
		flash := beego.NewFlash()
		flash.Success(email)
		flash.Store(&this.Controller)
		event.OnAccountCreated(email, nickname, sign_up_form.UserID) //todo
	} else {
		s := utils.NewInstant(errs, map[string]string{"email":  email, "password": ""})
		this.Data["json"] = &utils.SimpleJsonResponse{Status:0, Error:&s}
	}
	this.ServeJSON()
}

func (this *AccountController) SignUpSuccess() {
	flash := beego.ReadFromRequest(&this.Controller)
	if s, ok := flash.Data["success"]; ok {
		this.Data["email"] = s
		this.TplName = "account/signup_success.html"
	} else {
		this.Abort("404")
	}
}

func (this *AccountController) SignOut() {
	this.LogoutUser()
	this.Redirect("/account/signin", 302)
}