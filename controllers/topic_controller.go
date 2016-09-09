package controllers

import (
	"encoding/json"
	"gensh.me/goforum/models/database"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/context/posts"
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
	dbPosts := []m.Posts{}
	database.O.QueryTable("posts").Filter("visible", true).Limit(20, uint(start)).RelatedSel("Author").All(&dbPosts);
	mItems := posts.DBHotPostsConvert(&dbPosts)
	this.Data["slug"] = slug
	data, err := json.Marshal(mItems)
	if err == nil {
		this.Data["post_items"] = string(data)
		this.TplName = "topic/index.html"
		return
	}
	this.Abort("404")
}