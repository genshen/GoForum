package controllers

import (
	"gensh.me/goforum/models/database"
	"gensh.me/goforum/models/m"
	"encoding/json"
)

type TopicController struct {
	BaseController
}

var topic_rules = map[string]int{
}

func (this *TopicController) getRules(action string) int {
	return post_rules[action]
}

func (this *TopicController) Slug() {
	slug := this.Ctx.Input.Param("slug")
	start := 0
	dbPosts :=  []m.Posts{}
	database.DB.Where("visible = ?", true).Offset(uint(start)).Limit(20).Preload("Author").Preload("Author.Profile").Find(&dbPosts);
	mItems := DBHotPostsConvert(&dbPosts)
	this.Data["slug"] = slug
	json,err := json.Marshal(mItems)
	if err == nil {
		this.Data["post_items"] = string(json)
		this.TplName = "topic/index.html"
		return
	}
	this.Abort("404")
}
