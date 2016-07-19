package controllers

import (
	"strconv"
	"encoding/json"
	identify "./../verify/auth"
)

type ProfileController struct {
	BaseController
}

var profile_rules = map[string]int{
	"Login":   0,
	"Following": identify.Login,
	"Followed": identify.Login,
}

func (this *ProfileController) getRules(action string) int {
	return profile_rules[action]
}

func (this *ProfileController) Person() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":uid"))
	uid := uint(id)
	can_edit := false
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

/*those persons i'am focusing */
func (this *ProfileController) Following() {
	follows := findFollowsById(this.getUserId(),true)
	json, err := json.Marshal(follows)
	if err == nil {
		this.Data["follows"]  = string(json)
		this.TplName = "profile/following.html"
		return
	}
	this.Abort("404")
}

func (this *ProfileController) Followed() {
	follows := findFollowsById(this.getUserId(),false)
	json, err := json.Marshal(follows)
	if err == nil {
		this.Data["follows"]  = string(json)
		this.TplName = "profile/followed.html"
		return
	}
	this.Abort("404")
}

func (this *ProfileController) Collection() {
	this.TplName = "profile/collection.html"
}