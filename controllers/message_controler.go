package controllers

import (
	"gensh.me/goforum/components/context/message"
	"gensh.me/goforum/components/utils"
)

type MessageController struct {
	BaseController
}

var message_rules = map[string]int{
	"Message": utils.LoginJSON,
}

func (this *MessageController) getRules(action string) int {
	return message_rules[action]
}

func (this *MessageController) Messages(){
	mMessages := message.FindLatestPostMessages(this.getUserId(),[]int{utils.POST_COMMENT,utils.POST_REPLY})
	this.Data["json"] = &mMessages
	this.ServeJSON()
}

func (this *MessageController) Notifications(){
	mNotifications := message.FindLatestNotifications(this.getUserId(),[]int{utils.FOLLOW_ADD})
	this.Data["json"] = &mNotifications
	this.ServeJSON()
}