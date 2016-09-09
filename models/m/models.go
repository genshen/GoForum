package m

import (
	"time"
	"strconv"
	"gensh.me/goforum/models/database"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Follow))
	orm.RegisterModel(new(Posts), new(Comment), new(Topic), new(Tag))
	orm.RegisterModel(new(PostMessage), new(Notification))
	orm.RegisterModel(new(Swipe), new(Feedback))
	database.O = orm.NewOrm()
}


//User         User `gorm:"ForeignKey:Author"`
type Posts struct {
	Id           uint        `orm:"pk"`
	CreatedAt    time.Time   `orm:"auto_now_add"`
	UpdatedAt    time.Time   `orm:"auto_now_add"`
	DeletedAt    time.Time
			  //TopicId      uint
	Topic        *Topic  `orm:"rel(fk)"`
	Tags         []*Tag  `orm:"rel(m2m);rel_table(post_tag)"`
			  //AuthorId     uint
	Author       *Profile `orm:"rel(fk)"`
	Title        string   `orm:"size(255)"`
	Summary      string   `orm:"size(255)"`
	Content      string
	IsMobile     bool // 1 for mobile,0 for desktop
	Sticky       bool   `orm:"default(false)"`
	CommentCount int    `orm:"default(0)"`
	ViewCount    int    `orm:"default(0)"`
	Visible      bool   `orm:"default(true)"`
	LastReplayAt time.Time
			  //Comment      []Comment  `gorm:"ForeignKey:PostID"` //todo
}

func (this *Posts)TableName() string {
	return "posts"
}

func (p *Posts) GetPostById(id string) (err error) {
	//todo string to uint
	//database.O.One(&p, id)
	pid, _ := strconv.Atoi(id)
	p.Id = uint(pid)
	if err := database.O.Read(p); err == nil {
		_, e := database.O.LoadRelated(p, "Author")
		return e
	}
	return
}

func (p *Posts) Exist(id uint) bool {
	database.O.QueryTable("posts").Filter("id",id).Limit(1).
	One(p,"id", "author_id","summary","title","comment_count")
	//database.O.Select("id,author_id,summary,title,comment_count").First(&p, id)
	return p.Id != 0
}

type Comment struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now_add"`
	DeletedAt time.Time

	PostId    uint
	Author    *Profile `orm:"rel(fk)"`
	Parent    uint     `orm:"default(0)"`
	Content   string
	Visible   bool   `orm:"default(true)"`
}

func (c *Comment) TableName() string {
	return "comment"
}


type Topic struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now_add"`
	DeletedAt time.Time

	Name      string
	Describe  string
	Slug      string
	Visible   bool    `orm:"default(true)"`
	Color     string
}

func (this *Topic) TableName() string {
	return "topic"
}

type Tag struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now_add"`
	DeletedAt time.Time

	Name      string
	Describe  string
	Visible   bool   `orm:"default(true)"`
	Color     string
}

func (this *Tag) TableName() string {
	return "tag"
}

type Swipe struct {
	Id        uint        `orm:"pk"`
	CreatedAt time.Time   `orm:"auto_now_add"`
	UpdatedAt time.Time   `orm:"auto_now_add"`
	DeletedAt time.Time

	Url       string
	Img       string
	Content   string
	Visible   bool    `orm:"default(true)"`
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
	UpdatedAt time.Time   `orm:"auto_now_add"`
	DeletedAt time.Time

	UserId    uint
	Type      int8
	Content   string
	Contact   string
}

func (this *Feedback) TableName() string {
	return "feedback"
}
