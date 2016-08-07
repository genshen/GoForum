package profile

import (
	"gensh.me/goforum/models/database"
	"github.com/astaxie/beego/orm"
)

const (
	FollowerSQL = "select profile.id,profile.bio,profile.name,profile.avatar from follows " +
	"inner join profile on profile.id = follows.following_id where follower_id = ? limit 256"
	FollowingSQL = "select profile.id,profile.bio,profile.name,profile.avatar from follows " +
	"inner join profile on profile.id = follows.follower_id where following_id = ? limit 256"
)

type AllFollows struct {
	Following *[]orm.Params
	Followed  *[]orm.Params
}

func FindFollowsById(id uint) *AllFollows {
	//todo Limit(256).
	var follower_maps, following_maps  []orm.Params
	database.O.Raw(FollowerSQL, id).Values(&follower_maps)
	database.O.Raw(FollowingSQL, id).Values(&following_maps)
	//database.DB.Where("following_id = ?", id).Limit(256).Preload("Follower").Preload("Follower.Profile").Find(&db_followed)
	return &AllFollows{Following:&follower_maps, Followed:&following_maps}
}
