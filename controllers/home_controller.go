package controllers

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
	"gensh.me/goforum/components/context/posts"
	"gensh.me/goforum/components/utils"
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
	start := this.Ctx.Input.Param(":start")
	this.Data["json"] = posts.LoadHotPost(start)
	this.ServeJSON()
}

func (this *HomeController) Category() {
	mCategory :=  posts.Category{}
	mCategory.NewInstant()
	this.Data["json"] = &mCategory
	this.ServeJSON()
}

func (this *HomeController) Me() {
	var me utils.UserStatus
	if !this.IsUserLogin() {
		me.IsLogin = false
	} else {
		me.IsLogin = true
		me.ID = this.getUserId()
		me.Name = this.getUsername()
		profile := m.Profile{}  //load avatar info from database(while id username from session)
		database.O.QueryTable("profile").Filter("id", me.ID).Limit(1).One(&profile, "avatar")
		me.Avatar = profile.Avatar
	}
	this.Data["json"] = &me
	this.ServeJSON()
}