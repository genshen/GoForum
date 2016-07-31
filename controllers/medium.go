package controllers

import (
	"time"
	"../models/m"
	"../models/database"
)

type Person struct {
	ID     uint
	Name   string
	Avatar string
}

type Profile struct {
	m.Profile
	Name        string
	HasFollowed bool
	Edit        bool
	IsLogin     bool
}

func GetProfileById(my_id uint, uid uint) (profile Profile) {
	user := m.User{}
	database.DB.Preload("Profile").First(&user, uid)
	var follow_count int
	database.DB.Model(&m.Follow{}).Where("follower_id = ? AND following_id = ?", my_id,uid).Count(&follow_count)
	profile = Profile{Name:user.Name, Profile:user.Profile, Edit:(my_id == uid),
		HasFollowed:follow_count != 0, IsLogin:!(my_id == 0)}
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
			Person:Person{ID:db_hot.Author.ID, Name:db_hot.Author.Name, Avatar:db_hot.Author.Profile.Avatar}});
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
type AllFollows struct {
	Following *[]PersonFollow
	Followed  *[]PersonFollow
}
/*find all person who is following */
/*find all person who is followed by*/
func findFollowsById(id uint) *AllFollows {
	db_following := []m.Follow{}
	db_followed := []m.Follow{}
	//todo Limit(256).
	database.DB.Where("follower_id = ?", id).Limit(256).Preload("Following").Preload("Following.Profile").Find(&db_following)
	personFollowing := make([]PersonFollow, 0, len(db_following))
	for _, follow := range db_following {
		personFollowing = append(personFollowing, PersonFollow{Bio:follow.Following.Profile.Bio,
			Person:Person{ID:follow.Following.ID, Name:follow.Following.Name, Avatar:follow.Following.Profile.Avatar}});
	}
	//todo Limit(256)
	database.DB.Where("following_id = ?", id).Limit(256).Preload("Follower").Preload("Follower.Profile").Find(&db_followed)
	personFollowed := make([]PersonFollow, 0, len(db_followed))  //dbHotPosts to mHotPosts
	for _, follow := range db_followed {
		personFollowed = append(personFollowed, PersonFollow{Bio:follow.Follower.Profile.Bio,
			Person:Person{ID:follow.Follower.ID, Name:follow.Follower.Name, Avatar:follow.Follower.Profile.Avatar}});
	}
	return &AllFollows{Following:&personFollowing, Followed:&personFollowed}
}