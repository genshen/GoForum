package controllers

import (
	"../models/m"
	"../models/database"
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
	database.DB.Where("visible = ?", true).Offset(uint(start)).Limit(20).Preload("Author").Preload("Author.Profile").Find(&dbHotPosts);
	mHot := DBHotPostsConvert(&dbHotPosts)
	this.Data["json"] = mHot
	this.ServeJSON()
}

func (this *HomeController)Category() {
	mCategory := Category{}
	mCategory.NewInstant()
	this.Data["json"] = &mCategory
	this.ServeJSON()
}

func (this *HomeController)Me() {
	var me UserStatus
	if !this.IsUserLogin() {
		me.IsLogin = false
	} else {
		me.IsLogin = true
		me.ID = this.getUserId()
		me.Name = this.getUsername()
		profile := m.Profile{}  //load cover info from database(while id username from session)
		database.DB.Select("cover").Where("user_refer = ?",me.ID).First(&profile)
		me.Cover = profile.Cover
	}
	this.Data["json"] = &me
	this.ServeJSON()
}