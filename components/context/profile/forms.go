package profile

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/values"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/models/database"
)

type FollowAddForm struct {
	PersonID uint
}

func (this *FollowAddForm)Add(my_id uint) (result *utils.SimpleJsonResponse, isAdded bool) {
	isAdded = false
	if (my_id == this.PersonID) {
		result = &utils.SimpleJsonResponse{Status:0, Error:"不能关注自己本人哟~"}
		return
	}
	u := m.User{}
	database.O.QueryTable("user").Filter("id", this.PersonID).Filter("status", values.STATUS_ACTIVE).Limit(1).One(&u,"id") //todo FREEZING
	if u.Id == 0 {
		result = &utils.SimpleJsonResponse{Status:0, Error:"对应用户不存在"}
	} else {
		follow := m.Follow{}
		if database.O.QueryTable("follows").Filter("follower_id", my_id).Filter("following_id", this.PersonID).Limit(1).One(&follow);
		follow.Follower.Id == 0 && follow.Following.Id == 0 {
			database.O.Insert(&m.Follow{Follower:follow.Follower, Following:follow.Following})
			result = &utils.SimpleJsonResponse{Status:1, Addition:this.PersonID}
			isAdded = true
		} else {
			result = &utils.SimpleJsonResponse{Status:0, Error:"已经关注了改用户"}
		}
	}
	return
}