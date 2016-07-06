package controllers

import (
	"../models/m"
	"strconv"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.TplName = "home/index.html"
}

func (this *HomeController) LoadSwipe() {
	swipes_form := m.LoadSwipes()
	this.Data["json"] = swipes_form
	this.ServeJSON()
}

func (this *HomeController) Hot() {
	start,_ := strconv.Atoi(this.Ctx.Input.Param(":start"))
	this.Data["json"] = m.GetHotPosts(start)
	this.ServeJSON()
}