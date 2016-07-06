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

type HotPost struct {
	Person
	PostID       uint
	Title        string
	ViewCount    int
	CommentCount int
}

func (this *HomeController) Hot() {
	start, _ := strconv.Atoi(this.Ctx.Input.Param(":start"))
	dbHotPosts := m.GetHotPosts(start)
	mHotPosts := make([]HotPost, 0, len(dbHotPosts))  //m.Comments to Comments
	for _, db_hot := range dbHotPosts {
		mHotPosts = append(mHotPosts, HotPost{PostID:db_hot.ID, Title:db_hot.Title,
			ViewCount:db_hot.ViewCount, CommentCount:db_hot.CommentCount,
			Person:Person{ID:db_hot.Author.ID, Name:db_hot.Author.Name, Head:""}});
	}
	this.Data["json"] = &mHotPosts
	this.ServeJSON()
}