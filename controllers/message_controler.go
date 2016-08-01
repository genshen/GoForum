package controllers

import (
	identify "../middleware/values"
)

type MessageController struct {
	BaseController
}

var message_rules = map[string]int{
	"Message": identify.LoginJSON,
}

func (this *MessageController) getRules(action string) int {
	return message_rules[action]
}

func (this *MessageController) Messages(){
	mMessages := findLatestMessages(this.getUserId(),[]int{identify.POST_COMMENT,identify.POST_REPLY,})
	this.Data["json"] = &mMessages
	this.ServeJSON()
}

func (this *MessageController) Notifications(){
	mMessages := findLatestMessages(this.getUserId(),[]int{identify.FOLLOW_ADD})
	this.Data["json"] = &mMessages
	this.ServeJSON()
}