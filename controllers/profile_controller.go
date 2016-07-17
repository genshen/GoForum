package controllers

import (
	"strconv"
	"encoding/json"
	identify "./../verify/auth"
)

type ProfileController struct {
	BaseController
}

var peifile_rules = map[string]int{
	"Login":   0,
	"Logout": identify.Login,
}

func (this *ProfileController) getRules(action string) int {
	return peifile_rules[action]
}

func (this *ProfileController) Person() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":uid"))
	uid := uint(id)
	can_edit := false;
	if this.getUserId() == uid {
		//is self
		can_edit = true
	}
	profile := GetProfileById(uid, can_edit)
	if profile.Profile.UserRefer != 0 {
		json, err := json.Marshal(profile)
		if err == nil {
			this.Data["person"] = string(json)
			this.TplName = "profile/person.html"
			return
		}
	}
	this.Abort("404")
}

func (this *ProfileController) Following() {
	this.TplName = "profile/following.html"
}

func (this *ProfileController) Followed() {
	this.TplName = "profile/followed.html"
}

func (this *ProfileController) Collection() {
	this.TplName = "profile/collection.html"
}