package controllers

import (
	"strconv"
	"encoding/json"
	"gensh.me/goforum/components/context/profile"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/components/event"
	identify "gensh.me/goforum/models/values"
)

type ProfileController struct {
	BaseController
}

var profile_rules = map[string]int{
	"Follow": identify.Login | identify.JumpBack,
	"Collection": identify.Login | identify.JumpBack,
	"FollowAdd":identify.LoginJSON,
}

func (this *ProfileController) getRules(action string) int {
	return profile_rules[action]
}

func (this *ProfileController) Person() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":uid"))
	uid := uint(id)
	profile, exists := GetProfileById(this.getUserId(), uid)
	if exists {
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
		this.Data["follows"] = string(json)
		this.TplName = "profile/follow.html"
		return
	}
	this.Abort("404")
}

func (this *ProfileController) FollowAdd() {
	var result *utils.SimpleJsonResponse
	id, err := this.GetInt("id")
	if err == nil {
		faf := profile.FollowAddForm{PersonID:uint(id)}
		myID := this.getUserId()
		create_result, is_created := faf.Add(myID)
		result = create_result
		if is_created {
			event.OnFollowed(uint(id), myID, this.getUsername())
		}
	} else {
		result = &utils.SimpleJsonResponse{Status:0, Error:"ID不合法"}
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ProfileController) Collection() {
	this.TplName = "profile/collection.html"
}