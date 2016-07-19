package controllers

import (
	"time"
	"../models/m"
	"../models/database"
)

type Person struct {
	ID    uint
	Name  string
	Cover string
}

type Profile struct {
	m.Profile
	Name string
	Edit bool
}

func GetProfileById(uid uint, can_edit bool) (profile Profile) {
	user := m.User{}
	database.DB.Preload("Profile").First(&user, uid)
	profile = Profile{Name:user.Name, Profile:user.Profile, Edit:can_edit}
	return
}

//<post item >
type PostItem struct {
	Person
	PostID       uint
	Title        string
	ViewCount    int
	CommentCount int
}

func DBHotPostsConvert(dbHotPosts *[]m.Posts) (*[]PostItem) {
	postItems := make([]PostItem, 0, len(*dbHotPosts))  //dbHotPosts to mHotPosts
	for _, db_hot := range *dbHotPosts {
		postItems = append(postItems, PostItem{PostID:db_hot.ID, Title:db_hot.Title,
			ViewCount:db_hot.ViewCount, CommentCount:db_hot.CommentCount,
			Person:Person{ID:db_hot.Author.ID, Name:db_hot.Author.Name, Cover:db_hot.Author.Profile.Cover}});
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
	database.DB.Where("visible = ?", true).Find(&this.Topics)
	database.DB.Where("visible = ?", true).Find(&this.Tags)
}
// </structs for Category >
//  </struct for me >
type UserStatus struct {
	Person
	IsLogin bool
}
//< /struct for me page >

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

/*follow person*/
type PersonFollow struct {
	Person
	Bio string
}
/*find all person who is following (select_following:true)*/
/*find all person who is followed by(select_following:false)*/
func findFollowsById(id uint, select_following bool) *[]PersonFollow {
	db_follows := []m.Follow{}
	if select_following {
		database.DB.Where("follower_id = ?", id).Preload("Following").Preload("Following.Profile").Find(&db_follows)
		personFollows := make([]PersonFollow, 0, len(db_follows))  //dbHotPosts to mHotPosts
		for _, follow := range db_follows {
			personFollows = append(personFollows, PersonFollow{Bio:follow.Following.Profile.Bio,
				Person:Person{ID:follow.Following.ID, Name:follow.Following.Name, Cover:follow.Following.Profile.Cover}});
		}
		return &personFollows
	} else {
		database.DB.Where("following_id = ?", id).Preload("Follower").Preload("Follower.Profile").Find(&db_follows)
		personFollows := make([]PersonFollow, 0, len(db_follows))  //dbHotPosts to mHotPosts
		for _, follow := range db_follows {
			personFollows = append(personFollows, PersonFollow{Bio:follow.Follower.Profile.Bio,
				Person:Person{ID:follow.Follower.ID, Name:follow.Follower.Name, Cover:follow.Follower.Profile.Cover}});
		}
		return &personFollows
	}
}