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
	"Follow": identify.Login,
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

/*those persons i'am focusing or be followed */
func (this *ProfileController) Follow() {
	follows := findFollowsById(this.getUserId())
	json, err := json.Marshal(follows)
	if err == nil {
		this.Data["follows"]  = string(json)
		this.TplName = "profile/follow.html"
		return
	}
	this.Abort("404")
}

func (this *ProfileController) Collection() {
	this.TplName = "profile/collection.html"
}