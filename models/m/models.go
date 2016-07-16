package m

import (
	"github.com/jinzhu/gorm"
	"time"
	"./../database"
)

const (
	STATUS_ACTIVE int = 10
)

type User struct {
	gorm.Model
	Email      string
	Tel        string
	Name       string
	Password   string  `gorm:"column:password_hash"`
	AuthKey    string  `gorm:"column:auth_key"`
	ResetToken string  `gorm:"column:password_reset_token"`
	Status     int     `gorm:"default:10"`
	Profile    Profile `gorm:"ForeignKey:UserRefer"`
}

type Profile struct {
	//ID uint `gorm:"primary_key"`
	UserRefer uint
	Head      string `gorm:"default:'default.png'"`
	Coins     int    `gorm:"default:0"`
}

func (Profile) TableName() string {
	return "profile"
}

func (User) TableName() string {
	return "user"
}

func (u *User) GetUserById(id uint) {
	database.DB.Preload("Profile").First(&u, id)
}

//User         User `gorm:"ForeignKey:Author"`
type Posts struct {
	gorm.Model
	TopicID      uint
	Topic        Topic `gorm:"ForeignKey:TopicID"`
	Tags         []Tag `gorm:"many2many:post_tag;"`
	AuthorID     uint
	Author       User `gorm:"ForeignKey:AuthorID"`
	Title        string
	Content      string
	IsMobile     bool // 1 for mobile,0 for desktop
	Sticky       bool  `gorm:"default:false"`
	CommentCount int   `gorm:"default:0"`
	ViewCount    int   `gorm:"default:0"`
	Visible      bool  `gorm:"default:true"`
	LastReplayAt time.Time
	Comment      []Comment `gorm:"ForeignKey:PostID"`
}

func (p *Posts) GetPostById(id string) {
	//todo string to uint
	database.DB.First(&p, id)
}

func (p *Posts) Exist(id uint) bool {
	database.DB.Select("id,comment_count").First(&p, id)
	if p.ID == 0 {
		return false
	}
	return true
}

type Comment struct {
	gorm.Model
	PostID  uint
	Author  uint
	Parent  uint `gorm:"default=0"`
	Content string
	Visible bool  `gorm:"default:true"`
}

func (Comment) TableName() string {
	return "comment"
}

func LoadComments(id int, offset int) []Comment {
	var comments []Comment
	database.DB.Where("post_id = ?", id).Offset(uint(offset)).Limit(20).Find(&comments)
	return comments
}

type Topic struct {
	gorm.Model
	Name     string
	Describe string
	Slug     string
	Visible  bool  `gorm:"default:true"`
	Color    string
}

func (Topic) TableName() string {
	return "topic"
}

type Tag struct {
	gorm.Model
	Name     string
	Describe string
	Visible  bool  `gorm:"default:true"`
	Color    string
}

func (Tag) TableName() string {
	return "tag"
}

type Swipe struct {
	gorm.Model
	Url     string
	Img     string
	Content string
	Visible bool  `gorm:"default:true"`
}

func (Swipe) TableName() string {
	return "swipe"
}

func LoadSwipes() (swipes []Swipe) {
	database.DB.Where("visible = ?", true).Limit(8).Find(&swipes)
	return
}