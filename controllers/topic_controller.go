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
	start := c.Ctx.Input.Param(":start")
	topic_id := c.Ctx.Input.Param(":topic_id") //no need to int
	c.Data["json"] = posts.LoadPostsByTopicId(topic_id, start)
	c.ServeJSON()
}

func (this *TopicController) Slug() {
	slug := this.Ctx.Input.Param(":slug")
	topic := m.Topic{}
	err := database.O.QueryTable("topic").Filter("slug", slug).One(&topic) //todo remove deletedAt
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