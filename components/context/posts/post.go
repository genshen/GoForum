package posts

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/models/database"
	"time"
)

//<post item >
type PostItem struct {
	utils.Person
	PostID       uint
	Title        string
	ViewCount    int
	CommentCount int
}


/**<used for Post detail> */
type PostDetail struct {
	ID           uint
	//Topic
	Title        *string
	Content      *string
	IsMobile     bool
	Sticky       bool
	CommentCount int
	ViewCount    int
	LastReplayAt *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (this *PostDetail) NewInstant(p *m.Posts) {
	this.ID = p.Id
	this.Title = &p.Title
	this.Content = &p.Content
	this.IsMobile = p.IsMobile
	this.Sticky = p.Sticky
	this.CommentCount = p.CommentCount
	this.ViewCount = p.ViewCount
	this.LastReplayAt = &p.LastReplayAt
	this.CreatedAt = &p.CreatedAt
	this.UpdatedAt = &p.UpdatedAt
}

type PostView struct {
	IsLogin bool
	Post    PostDetail
	Author  utils.Person
}
/*</used for Post detail> */

func DBHotPostsConvert(dbHotPosts *[]m.Posts) (*[]PostItem) {
	postItems := make([]PostItem, 0, len(*dbHotPosts))  //dbHotPosts to mHotPosts
	for _, db_hot := range *dbHotPosts {
		postItems = append(postItems, PostItem{PostID:db_hot.Id, Title:db_hot.Title,
			ViewCount:db_hot.ViewCount, CommentCount:db_hot.CommentCount,
			Person:utils.Person{ID:db_hot.Author.Id, Name:db_hot.Author.Name, Avatar:db_hot.Author.Avatar}});
	}
	return &postItems
}
//</post item >


// <structs for Category >
//type Tag struct {
//
//}

//type Topic struct {
//
//}

type Category struct {
	Tags   []m.Tag
	Topics []m.Topic
}

func (this *Category)NewInstant() {
	database.O.QueryTable("topic").Filter("visible", true).All(&this.Topics)
	database.O.QueryTable("tag").Filter("visible", true).All(&this.Tags)
	//RW database.DB.Where("visible = ?", true).Find(&this.Topics)
	//RW database.DB.Where("visible = ?", true).Find(&this.Tags)
}
