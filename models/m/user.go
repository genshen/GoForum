package m

import (
	"gensh.me/goforum/models/database"
	"time"
)

type User struct {
	Id         uint        `orm:"pk"`
	CreatedAt  time.Time   `orm:"auto_now_add"`
	UpdatedAt  time.Time   `orm:"auto_now_add"`
	DeletedAt  time.Time

	Email      string
	Tel        string
	Name       string
	Password   string   `orm:"column(password_hash)"`
	AuthKey    string   `orm:"column(auth_key)"`
	ResetToken string   `orm:"column(password_reset_token)"`
	Status     int      `orm:"default(1)"` //todo UNACTIVATED
	Profile    *Profile `orm:"reverse(one)"`
}

type Profile struct {
	Id             uint    `orm:"pk"`
	UserRefer      *User   `orm:"rel(one)"` //UserReferId is total the same as Id
	Name           string  `orm:"default('')"`
	Avatar         string  `orm:"default('default.png')"`
	Coins          int     `orm:"default(0)"`
	PostCount      int     `orm:"default(0)"`
	CommentCount   int     `orm:"default(0)"`
	FollowedCount  int     `orm:"default(0)"`
	FollowingCount int     `orm:"default(0)"`
	Bio            string  `orm:"default('这家伙很懒,什么都没有')"`
}

func (this *Profile) TableName() string {
	return "profile"
}

func (this *User) TableName() string {
	return "user"
}

func (u *User) GetUserById(id uint) {
	//database.DB.Preload("Profile").First(&u, id)
	database.O.QueryTable("user").Filter("id", id).Limit(1).One(u)
}

type Follow struct {
	Follower  *Profile  `orm:"rel(fk)"`
	Following *Profile  `orm:"rel(fk)"`
	Id        uint      `orm:"pk"`
}

func (this *Follow) TableName() string {
	return "follows"
}