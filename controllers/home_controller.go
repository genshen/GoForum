package controllers

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
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
	start, _ := strconv.Atoi(this.Ctx.Input.Param(":start"))
	dbHotPosts := []m.Posts{}
	//todo Preload("Author.Profile")
	database.O.QueryTable("posts").Filter("visible", true).Limit(20, uint(start)).
	RelatedSel("Author__Profile").All(&dbHotPosts);
	mHot := DBHotPostsConvert(&dbHotPosts)
	this.Data["json"] = mHot
	this.ServeJSON()
}

func (this *HomeController) Category() {
	mCategory := Category{}
	mCategory.NewInstant()
	this.Data["json"] = &mCategory
	this.ServeJSON()
}

func (this *HomeController) Me() {
	var me UserStatus
	if !this.IsUserLogin() {
		me.IsLogin = false
	} else {
		me.IsLogin = true
		me.ID = this.getUserId()
		me.Name = this.getUsername()
		profile := m.Profile{}  //load avatar info from database(while id username from session)
		database.O.QueryTable("profile").Filter("user_refer", me.ID).Limit(1).One(&profile, "avatar")
		me.Avatar = profile.Avatar
	}
	this.Data["json"] = &me
	this.ServeJSON()
}