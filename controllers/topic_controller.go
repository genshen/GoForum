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

func (c *TopicController) Posts() {
	//start := c.Ctx.Input.Param("slug")
	dbPosts := []m.Posts{}
	database.O.QueryTable("posts").Filter("visible", true).Limit(20, 20).RelatedSel("Author").All(&dbPosts) //todo errors
	mItems := posts.DBHotPostsConvert(&dbPosts)
	c.Data["json"] = mItems
	c.ServeJSON()
}
func (this *TopicController) Slug() {
	slug := this.Ctx.Input.Param(":slug")
	topic := m.Topic{}
	err := database.O.QueryTable("topic").Filter("slug", slug).One(&topic)
	if err == nil {
		data, err_ := json.Marshal(&topic)
		if err_ == nil {
			this.Data["topic_detail"] = string(data)
			this.TplName = "topic/index.html"
			return
		}
	}
	this.Abort("404")
}