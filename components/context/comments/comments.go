package comments

import (
	"gensh.me/goforum/models/database"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

var (
	load_comments_sql string
)

func init() {
	qb, _ := orm.NewQueryBuilder(beego.AppConfig.String("db_type"))
	qb.Select("comment.id", "comment.author_id", "comment.content", "comment.created_at",
		"profile.name", "profile.avatar").
		From("comment").
		InnerJoin("profile").On("comment.author_id = profile.id").
		Where("post_id = ? And visible = true").
	//OrderBy("id").Dec()
		Limit(20)
	load_comments_sql = qb.String()
}

func LoadComments(id string, offset string) *[]orm.Params {
	var comments []orm.Params
	database.O.Raw(load_comments_sql + " offset " + offset, id).Values(&comments)
	return &comments
}
