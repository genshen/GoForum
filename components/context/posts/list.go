package posts

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/models/database"
)

var (
	hot_post_sql string
	post_by_topic_id_sql string
)

func init() {
	qb, _ := orm.NewQueryBuilder(beego.AppConfig.String("db_type"))
	qb.Select("posts.id", "posts.author_id", "posts.title", "posts.comment_count", "posts.view_count", "posts.created_at",
		"profile.name", "profile.avatar").
		From("posts").
		InnerJoin("profile").On("posts.author_id = profile.id").
		Where("visible = true").
	//
		Limit(20)
	hot_post_sql = qb.String()

	qb_t, _ := orm.NewQueryBuilder(beego.AppConfig.String("db_type"))
	qb_t.Select("posts.id", "posts.author_id", "posts.title", "posts.sticky", "posts.comment_count",
		"posts.view_count", "posts.created_at",
		"profile.name", "profile.avatar").
		From("posts").
		InnerJoin("profile").On("posts.author_id = profile.id").
		Where("visible = true").
		And("topic_id = ?").
		OrderBy("sticky").Desc().
		Limit(20)
	post_by_topic_id_sql = qb_t.String()
}

//<post item >
type PostItem struct {
	utils.Person
	PostID       uint
	Title        string
	ViewCount    int
	CommentCount int
	CreatedAt    *time.Time
}

//</post item >
func DBHotPostsConvert(dbHotPosts *[]m.Posts) (*[]PostItem) {
	postItems := make([]PostItem, 0, len(*dbHotPosts))  //dbHotPosts to mHotPosts
	for _, db_hot := range *dbHotPosts {
		postItems = append(postItems, PostItem{PostID:db_hot.Id, Title:db_hot.Title,
			ViewCount:db_hot.ViewCount, CommentCount:db_hot.CommentCount, CreatedAt:&db_hot.CreatedAt,
			Person:utils.Person{ID:db_hot.Author.Id, Name:db_hot.Author.Name, Avatar:db_hot.Author.Avatar}});
	}
	return &postItems
}

func LoadHotPost(start string) *[]orm.Params {
	var maps []orm.Params
	database.O.Raw(hot_post_sql + " offset " + start).Values(&maps)
	return &maps
}

func LoadPostsByTopicId(id string,start string)*[]orm.Params {
	var maps []orm.Params
	database.O.Raw(post_by_topic_id_sql + " offset " + start,id).Values(&maps)
	return &maps
}