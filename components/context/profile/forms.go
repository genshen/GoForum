package profile

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/models/database"
	models_utils "gensh.me/goforum/models/utils"
)

const (
	InsertFollowSQL = "insert into follows values(?,?)"
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
	database.O.QueryTable("user").Filter("id", this.PersonID).Filter("status", utils.STATUS_ACTIVE).Limit(1).One(&u, "id") //todo FREEZING
	if u.Id == 0 {
		result = &utils.SimpleJsonResponse{Status:0, Error:"对应用户不存在"}
	} else {
		if models_utils.Counts("follows", "follower_id = ? and following_id = ?", my_id, this.PersonID) == 0 {
			_, err := database.O.Raw(InsertFollowSQL, my_id, this.PersonID).Exec()
			if err == nil {
				result = &utils.SimpleJsonResponse{Status:1, Addition:this.PersonID}
				isAdded = true
			} else {
				result = &utils.SimpleJsonResponse{Status:0, Error:"关注失败,请重试"}
			}
		} else {
			result = &utils.SimpleJsonResponse{Status:0, Error:"已经关注了改用户"}
		}
	}
	return
}