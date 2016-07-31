package controllers

import (
	identify "../middleware/values"
)

type MessageController struct {
	BaseController
}

const message_rules = map[string]int{
	"Message": identify.LoginJSON,
}

func (this *MessageController) getRules(action string) int {
	return message_rules[action]
}

func (this *MessageController) Message(){
	mMessages := findLatestMessages(this.getUserId())
	this.Data["json"] = &mMessages
	this.ServeJSON()
}