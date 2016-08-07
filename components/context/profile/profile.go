package profile

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
	models_utils "gensh.me/goforum/models/utils"
)

type Profile struct {
	*m.Profile
	HasFollowed bool
	Edit        bool
	IsLogin     bool
}

func GetProfileById(my_id uint, uid uint) (profile Profile, exists bool) {
	exists = false
	_profile := m.Profile{}
	database.O.QueryTable("profile").Filter("id", uid).Limit(1).One(&_profile)
	if _profile.Id == 0 {
		return
	}
	_profile.UserRefer = nil

	follow_count := models_utils.Counts("follows", "follower_id = ? and following_id = ?", my_id, uid)
	profile = Profile{Profile:&_profile, Edit:(my_id == uid), HasFollowed:follow_count != 0, IsLogin:!(my_id == 0)}
	exists = true
	return
}

