package m

import (
	"gensh.me/goforum/models/database"
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Follow))
	orm.RegisterModel(new(Posts), new(Comment), new(Topic), new(Tag))
	orm.RegisterModel(new(PostMessage), new(Notification))
	orm.RegisterModel(new(Swipe),new(Feedback))
	database.O = orm.NewOrm()
}


//User         User `gorm:"ForeignKey:Author"`
type Posts struct {
	Id           uint        `orm:"pk"`
	CreatedAt    time.Time   `orm:"auto_now_add"`
	UpdatedAt    time.Time   `orm:"auto_now"`
	DeletedAt    time.Time

	//TopicId      uint
	Topic        *Topic  `orm:"rel(fk)"`
	Tags         []*Tag  `orm:"rel(m2m);rel_table(post_tag)"`
	//AuthorId     uint
	Author       *User   `orm:"rel(fk)"`
	Title        string  `orm:"size(255)"`
	Summary      string  `orm:"size(255)"`
	Content      string
	IsMobile     bool // 1 for mobile,0 for desktop
	Sticky       bool   `orm:"default(false)"`
	CommentCount int    `orm:"default(0)"`
	ViewCount    int    `orm:"default(0)"`
	Visible      bool   `orm:"default(true)"`
	LastReplayAt time.Time
			  //Comment      []Comment  `gorm:"ForeignKey:PostID"` //todo
}

func (this *Posts)TableName() string{
	return "posts"
}

func (p *Posts) GetPostById(id string) error {
	//todo string to uint
	//database.O.One(&p, id)
	return database.O.Read(&p);
}

func (p *Posts) Exist(id uint) bool {
	database.O.Raw("SELECT id, author_id,summary,title,comment_count FROM user WHERE id = ? LIMIT 1", id).QueryRow(&p)
	//database.O.Select("id,author_id,summary,title,comment_count").First(&p, id)
	return p.Id != 0
}

type Comment struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now"`
	DeletedAt time.Time

	PostID  uint
	Author  uint
	Parent  uint   `orm:"default(0)"`
	Content string
	Visible bool   `orm:"default(true)"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func LoadComments(id int, offset int) []Comment {
	var comments []Comment
	database.O.QueryTable("comment").Filter("post_id", id).Offset(uint(offset)).Limit(20).All(&comments)
	//database.DB.Where("post_id = ?", id).Offset(uint(offset)).Limit(20).Find(&comments)
	return comments
}

type Topic struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now"`
	DeletedAt time.Time

	Name     string
	Describe string
	Slug     string
	Visible  bool    `orm:"default(true)"`
	Color    string
}

func (this *Topic) TableName() string {
	return "topic"
}

type Tag struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now"`
	DeletedAt time.Time

	Name     string
	Describe string
	Visible  bool   `orm:"default(true)"`
	Color    string
}

func (this *Tag) TableName() string {
	return "tag"
}

type Swipe struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now"`
	DeletedAt time.Time

	Url     string
	Img     string
	Content string
	Visible bool    `orm:"default(true)"`
}

func (this *Swipe) TableName() string {
	return "swipe"
}

func LoadSwipes() (swipes []Swipe) {
	database.O.QueryTable("swipe").Filter("visible", true).Limit(20).All(&swipes)
	//database.DB.Where("visible = ?", true).Limit(8).Find(&swipes)
	return
}

type Feedback struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now"`
	DeletedAt time.Time

	UserID  uint
	Type    int8
	Content string
	Contact string
}

func (this *Feedback) TableName() string {
	return "feedback"
}
