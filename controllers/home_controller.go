package controllers

import ()

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "home/index.html"
}

type  SwipeData struct {

}

func (this *HomeController) LoadSwipe(){
	this.Data["json"] =
	this.ServeJSON()
}