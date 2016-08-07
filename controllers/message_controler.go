package controllers

import (
	identify "gensh.me/goforum/models/values"
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
	mMessages := findLatestPostMessages(this.getUserId(),[]int{identify.POST_COMMENT,identify.POST_REPLY})
	this.Data["json"] = &mMessages
	this.ServeJSON()
}

func (this *MessageController) Notifications(){
	mNotifications := findLatestNotifications(this.getUserId(),[]int{identify.FOLLOW_ADD})
	this.Data["json"] = &mNotifications
	this.ServeJSON()
}