package controllers

import (
	identify "../middleware/values"
)

type MessageController struct {
	BaseController
}

var message_rules = map[string]int{
	"Follow": identify.Login | identify.JumpBack,
	"Collection": identify.Login | identify.JumpBack,
}

func (this *MessageController) getRules(action string) int {
	return message_rules[action]
}

func (this *MessageController) Message(){
	mMessages := findLatestMessages(this.getUserId())
	this.Data["json"] = &mMessages
	this.ServeJSON()
}