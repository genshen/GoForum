package controllers

import (
	"time"
	"../models/m"
	"../models/database"
)

type Person struct {
	ID   uint
	Name string
	Head string
}

//< hot post >
type HotPost struct {
	Person
	PostID       uint
	Title        string
	ViewCount    int
	CommentCount int
}
//</ hot post >


// <structs for Category >
type Tag struct {

}

//type Topic struct {
//
//}

type Category struct {
	Tags   []m.Tag
	Topics []m.Topic
}

func (this *Category)NewInstant() {
	database.DB.Where("visible = ?", true).Find(&this.Topics)
	database.DB.Where("visible = ?", true).Find(&this.Tags)
}
// </structs for Category >

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
	this.ID = p.ID
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
	Author  Person
}
/*</used for Post detail> */