package m

import (
	"github.com/jinzhu/gorm"
	"time"
	"gensh.me/goforum/models/database"
)

//User         User `gorm:"ForeignKey:Author"`
type Posts struct {
	gorm.Model
	TopicID      uint
	Topic        Topic `gorm:"ForeignKey:TopicID"`
	Tags         []Tag `gorm:"many2many:post_tag;"`
	AuthorID     uint
	Author       User `gorm:"ForeignKey:AuthorID"`
	Title        string `gorm:"size:255"`
	Summary      string `gorm:"size:255"`
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
	database.DB.Select("id,author_id,summary,title,comment_count").First(&p, id)
	return p.ID != 0
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

type Feedback struct {
	gorm.Model
	UserID  uint
	Type    int8
	Content string
	Contact string
}

func (Feedback) TableName() string {
	return "feedback"
}
